package checkout

import (
	"fmt"
	"math"

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
		var itemSubTotal int
		var err error

		if rule.Offer != nil && rule.Offer.Quantity > 0 && quantity >= rule.Offer.Quantity {
			itemSubTotal, err = c.calculateBundles(quantity, rule)
			if err != nil {
				return 0, fmt.Errorf("error calculating price for SKU %s: %v", sku, err)
			}

		} else {
			if rule.UnitPrice > 0 && quantity > math.MaxInt/rule.UnitPrice {
				return 0, fmt.Errorf("price overflow for SKU %s", sku)
			}
			itemSubTotal = quantity * rule.UnitPrice
		}

		if itemSubTotal > math.MaxInt-totalPrice {
			return 0, fmt.Errorf("price overflow while adding SKU %s", sku)
		}
		totalPrice += itemSubTotal
	}

	return totalPrice, nil
}

// Calculates pricing for bundles that are eligible for special offers
func (c *checkout) calculateBundles(quantity int, pricingRule pricing.Rule) (int, error) {
	numOffers := quantity / pricingRule.Offer.Quantity
	remainingItems := quantity % pricingRule.Offer.Quantity

	if pricingRule.Offer.Price > 0 && numOffers > math.MaxInt/pricingRule.Offer.Price {
		return 0, fmt.Errorf("overflow on special offer price")
	}
	offerTotal := numOffers * pricingRule.Offer.Price

	if pricingRule.UnitPrice > 0 && remainingItems > math.MaxInt/pricingRule.UnitPrice {
		return 0, fmt.Errorf("overflow on remainder price")
	}

	remainderTotal := remainingItems * pricingRule.UnitPrice

	if offerTotal > math.MaxInt-remainderTotal {
		return 0, fmt.Errorf("overflow when summing offer and remainder prices")
	}

	return offerTotal + remainderTotal, nil
}
