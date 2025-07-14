package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegenem0/tm-tech-takehome/internal/pricing"
)

func setupPricingRules() pricing.Rules {
	return pricing.Rules{
		"A": {UnitPrice: 50},
		"B": {UnitPrice: 30},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
}

func TestGetTotalPrice_EmptyBasket(t *testing.T) {
	// setup
	checkout := NewCheckout(setupPricingRules())

	// act
	total, err := checkout.GetTotalPrice()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 0, total)
}

func TestGetTotalPrice_SingleBasicItem(t *testing.T) {

	// setup
	pricingRules := setupPricingRules()

	// act
	checkout := NewCheckout(pricingRules)

	err := checkout.Scan("C")
	assert.NoError(t, err)

	// assert

	total, err := checkout.GetTotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 20, total)
}
