package main

import (

)

// placeholder until we set up wrapper
var db = mockDB{}

type checkout struct {
	Basket map[string]int
}

func GetCheckout() checkout {
	return checkout{
		Basket: make(map[string]int),
	}
}

func (c checkout) scan(SKU string) error {
	// check that the product exists
	if _, err := db.getProduct(SKU); err != nil {
		return err
	}

	c.Basket[SKU]++
	return nil
}

func (c checkout) getTotalPrice() int {
	// todo
	return 1
}