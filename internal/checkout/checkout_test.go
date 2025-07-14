package checkout

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTotalPrice_EmptyBasket(t *testing.T) {
	// setup
	checkout := NewCheckout()

	// call
	total, err := checkout.GetTotalPrice()

	// assert
	assert.Nil(t, err)
	assert.Equal(t, 0, total)
}
