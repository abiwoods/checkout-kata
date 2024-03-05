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
	c := setupCheckout(t, products)

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
			input:    []string{"D"},
			expected: 15,
		},
		"basic - multiple of 1": {
			input:    []string{"C", "C", "C", "C"},
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
		"offer only": {
			input:    []string{"A", "A", "A"},
			expected: 130,
		},
		"offer and another item": {
			input:    []string{"A", "B", "A", "A"},
			expected: 130 + 30,
		},
		"offer and more of same item": {
			input:    []string{"A", "A", "A", "A", "A"},
			expected: 130 + (2 * 50),
		},
		"multiples of one offer": {
			input:    []string{"B", "B", "B", "B"},
			expected: 2 * 45,
		},
		"mixture of everything": {
			input:    []string{"D", "B", "C", "A", "A", "D", "A", "B", "B", "B", "A", "C", "D"},
			expected: 130 + 50 + (2 * 45) + (2 * 20) + (3 * 15),
		},
	}

	for name, test := range tests {
		if name == "mixture of everything" {
			assert.True(t, true)
		}

		c := setupCheckout(t, test.input)

		price := c.getTotalPrice()
		assert.Equal(t, test.expected, price, name)
	}
}

func setupCheckout(t *testing.T, SKUs []string) Checkout {
	c := GetCheckout()
	for _, s := range SKUs {
		err := c.scan(s)
		assert.NoError(t, err)
	}
	return c
}
