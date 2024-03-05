package main

import "fmt"

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
	total := 0
	for sku, count := range c.Basket {
		total += getProductTotal(sku, count)
	}
	
	return total
}

func getProductTotal(sku string, count int) int {
	product, err := db.getProduct(sku)
	if err != nil {
		fmt.Printf("tried to get details for %s - not found", sku)
		return 0
	}

	offerPrice, remainingCount := getOfferTotal(sku, count) 

	return offerPrice + product.UnitPrice * remainingCount
}

func getOfferTotal(sku string, count int) (int, int) {
	offer, err := db.getOffer(sku)
	if err != nil {
		return 0, count
	}

	offerCount := count/offer.Count
	remainingCount := count - offerCount * offer.Count
	offerTotal := offerCount * offer.Price

	return offerTotal, remainingCount
}