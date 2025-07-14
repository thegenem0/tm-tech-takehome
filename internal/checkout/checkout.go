package checkout

import (
	"fmt"

	"github.com/thegenem0/tm-tech-takehome/internal/pricing"
)

type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

type checkout struct {
	rules        pricing.Rules
	scannedItems map[string]int // maps SKU -> quantity
}

func NewCheckout(pricingRules pricing.Rules) ICheckout {
	return &checkout{
		rules:        pricingRules,
		scannedItems: make(map[string]int),
	}
}

func (c *checkout) Scan(SKU string) (err error) {
	if _, ok := c.rules[SKU]; !ok {
		return fmt.Errorf("item with SKU: %s does not exist in the store", SKU)
	}

	c.scannedItems[SKU]++

	return nil
}

func (c *checkout) GetTotalPrice() (totalPrice int, err error) {
	for sku, quantity := range c.scannedItems {
		rule := c.rules[sku]

		if rule.Offer != nil && quantity >= rule.Offer.Quantity {
			totalPrice += c.calculateBundles(quantity, rule)
		} else {
			totalPrice += quantity * rule.UnitPrice
		}
	}

	return totalPrice, nil
}

// Calculates pricing for bundles that are eligible for special offers
func (c *checkout) calculateBundles(quantity int, pricingRule pricing.Rule) int {
	numOffers := quantity / pricingRule.Offer.Quantity
	totalPrice := numOffers * pricingRule.Offer.Price

	remainingItems := quantity % pricingRule.Offer.Quantity
	totalPrice += remainingItems * pricingRule.UnitPrice

	return totalPrice
}
