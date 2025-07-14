package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalPrice_EmptyBasket(t *testing.T) {
	// TODO(thegenem0):
	// Implement constructors

	// pricingRules := setupPricingRules()
	// checkout := NewCheckout(pricingRules)

	var checkout ICheckout

	// TODO(thegenem0):
	// Implement Iface methods on Checkout
	total, err := checkout.GetTotalPrice()

	// Should not error
	assert.Nil(t, err)

	// Should be 0 for empty basket
	assert.Equal(t, 0, total)
}
