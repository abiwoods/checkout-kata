package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCheckoutReturnsEpmtyCheckoutStruct(t *testing.T) {
	c := GetCheckout()
	assert.Len(t, c.Basket, 0)
}

func TestScanIncreasesBasketCounts(t *testing.T) {
	products := []string{"A", "B", "C", "A"}
	c := setupCheckout(products)

	assert.Equal(t, 2, c.Basket["A"]) // checking order that these are input does not cause issues
	assert.Equal(t, 1, c.Basket["B"])
	assert.Equal(t, 1, c.Basket["C"])
	assert.Equal(t, 0, c.Basket["D"])

	c.scan("A")
	assert.Equal(t, 3, c.Basket["A"])
}

func TestScanErrorsForUnkownProduct(t *testing.T) {
	c := GetCheckout()
	err := c.scan("missing")
	assert.Error(t, err)
}

func TestGetTotalPrice(t *testing.T) {
	tests := map[string]struct {
		input    []string
		expected int
	}{
		"basic - one product": {
			input: []string{"D"},
			expected: 15,
		},
		"basic - multiple of 1": {
			input: []string{"C", "C", "C", "C"},
			expected: 4 * 20,
		},
		"basic - mixture": {
			input:    []string{"A", "B", "C", "C", "A", "D"},
			expected: 50 + 30 + 20 + 20 + 50 + 15,
		},
		"no products scanned": {
			input:    []string{},
			expected: 0,
		},
	}

	for _, test := range tests {
		c := setupCheckout(test.input)

		price := c.getTotalPrice()
		assert.Equal(t, test.expected, price)
	}
}

func setupCheckout(SKUs []string) checkout {
	c := GetCheckout()
	for _, s := range SKUs {
		c.scan(s)
	}
	return c
}
