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
	c := GetCheckout()

	c.scan("A")
	c.scan("B")
	c.scan("C")
	c.scan("A")

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

