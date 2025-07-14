package checkout

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thegenem0/tm-tech-takehome/internal/pricing"
)

func setupPricingRules() pricing.Rules {
	return pricing.Rules{
		"A": {UnitPrice: 50, Offer: nil},
		"B": {UnitPrice: 30, Offer: nil},
		"C": {UnitPrice: 20, Offer: nil},
		"D": {UnitPrice: 15, Offer: nil},
	}
}

func setupPricingRulesWithOffers() pricing.Rules {
	return pricing.Rules{
		"A": {UnitPrice: 50, Offer: &pricing.SpecialOffer{Quantity: 3, Price: 130}},
		"B": {UnitPrice: 30, Offer: &pricing.SpecialOffer{Quantity: 2, Price: 45}},
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

// Tests simple basket with no offers
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

// Tests basket with single offer
func TestGetTotalPrice_WithSpecialOffer(t *testing.T) {
	// setup
	pricingRules := setupPricingRulesWithOffers()
	checkout := NewCheckout(pricingRules)

	// act
	err := checkout.Scan("A")
	assert.NoError(t, err)

	err = checkout.Scan("A")
	assert.NoError(t, err)

	err = checkout.Scan("A")
	assert.NoError(t, err)

	total, err := checkout.GetTotalPrice()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 130, total)
}

// Tests complex basket with multiple items and offers
func TestGetTotalPrice_MixedBasketComplex(t *testing.T) {
	// setup
	checkout := NewCheckout(setupPricingRulesWithOffers())

	// act

	items := []string{"A", "B", "A", "B", "A", "A", "B", "A"}
	for _, item := range items {
		err := checkout.Scan(item)
		assert.NoError(t, err)
	}

	total, err := checkout.GetTotalPrice()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, 305, total)
}

func TestGetTotalPrice_IntOverflow(t *testing.T) {
	// setup
	pricingRules := pricing.Rules{
		"A": {UnitPrice: math.MaxInt / 2},
	}

	checkout := NewCheckout(pricingRules)

	// act

	err := checkout.Scan("A")
	assert.NoError(t, err)

	err = checkout.Scan("A")
	assert.NoError(t, err)

	err = checkout.Scan("A") // Third scan should trigger int overflow
	assert.NoError(t, err)

	// assert

	total, err := checkout.GetTotalPrice()
	assert.NoError(t, err)

	assert.Equal(t, math.MaxInt, total)
}

func TestGetTotalPrice_OfferWithZeroQuantity(t *testing.T) {
	// setup
	pricingRules := pricing.Rules{
		"A": {UnitPrice: 50, Offer: &pricing.SpecialOffer{Quantity: 0, Price: 100}},
	}
	checkout := NewCheckout(pricingRules)

	// act
	err := checkout.Scan("A")
	assert.NoError(t, err)

	// assert
	assert.NotPanics(t, func() {
		total, err := checkout.GetTotalPrice()
		assert.NoError(t, err)

		assert.Equal(t, 50, total)
	})
}

// Test for items that meet the exact eligibility boundary for offers
func TestGetTotalPrice_MultipleExactOffers(t *testing.T) {
	// setup
	checkout := NewCheckout(setupPricingRulesWithOffers())

	// act
	for range 6 {
		err := checkout.Scan("A")
		assert.NoError(t, err)
	}

	// assert
	total, err := checkout.GetTotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 260, total)
}

// Tests invalid SKU appropriately return error
func TestScan_InvalidSKU(t *testing.T) {
	// setup
	pricingRules := setupPricingRules()
	checkout := NewCheckout(pricingRules)

	// act
	err := checkout.Scan("X")

	// assert
	assert.Error(t, err)
}
