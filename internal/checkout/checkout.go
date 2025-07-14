package checkout

import "github.com/thegenem0/tm-tech-takehome/internal/pricing"

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
	c.scannedItems[SKU]++
	return nil
}

func (c *checkout) GetTotalPrice() (totalPrice int, err error) {
	for sku, quantity := range c.scannedItems {
		rule := c.rules[sku]
		totalPrice += quantity * rule.UnitPrice
	}

	return totalPrice, nil
}
