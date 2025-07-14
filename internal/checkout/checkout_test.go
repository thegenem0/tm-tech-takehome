package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalPrice_EmptyBasket(t *testing.T) {
	// setup
	checkout := NewCheckout()

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
