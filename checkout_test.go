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
	products := []string{ "A", "B", "C", "A"}
	c := setupCheckout(products)

	assert.Equal(t, 2, c.Basket["A"]) // checking order that these are input does not cause issues
	assert.Equal(t, 1, c.Basket["B"])
	assert.Equal(t, 1, c.Basket["C"])
	assert.Equal(t, 0, c.Basket["D"])
}

func TestScanErrorsForUnkownProduct(t *testing.T) {
	c := GetCheckout()
	err := c.scan("missing")
	assert.Error(t, err)
}

func TestGetTotalPriceBasic(t *testing.T) {
	c := setupCheckout([]string{"A", "B", "C", "C", "A", "D"})
	
	price := c.getTotalPrice()
	expected := 50 + 30 + 20 + 20 + 50 + 15
	assert.Equal(t, expected, price)
}

func setupCheckout(SKUs []string) checkout {
	c := GetCheckout()
	for _, s := range SKUs {
		c.scan(s)
	}
	return c
}

