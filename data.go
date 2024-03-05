package main

import (
	"fmt"
)

type dataSource interface {
	getProduct(string) (Product, error)
	getOffer(string) (Offer, error)
	setProductPrice(string, int)
	removeProduct(string)
	setOffer(string, int, int)
	removeOffer(string)
}

var products =  map[string]Product{
	"A": {
		SKU: "A",
		UnitPrice: 50,
	},
	"B": {
		SKU: "B",
		UnitPrice: 30,
	},
	"C": {
		SKU: "C",
		UnitPrice: 20,
	},
	"D": {
		SKU: "D",
		UnitPrice: 15,
	},
}

var offers = map[string]Offer{
	"A": {
		SKU: "A",
		Count: 3,
		Price: 130,
	},
	"B": Offer{
		SKU:"B",
		Count: 2,
		Price: 45,
	},
}

type mockDB struct {}

func (mockDB) getProduct(sku string) (Product, error) {
	result, ok := products[sku]

	if !ok {
		return Product{}, fmt.Errorf("product %s not found", sku)
	}

	return result, nil
}

func (mockDB) getOffer(sku string) (Offer, error) {
	result, ok := offers[sku]

	if !ok {
		return Offer{}, fmt.Errorf("offers for product %s not found", sku)
	}

	return result, nil
}

func (mockDB) setProductPrice(product Product) {
	products[product.SKU] = product
}

func (mockDB) removeProduct(sku string) {
	delete(products, sku)
}

func (mockDB) setOffer(offer Offer) {
	offers[offer.SKU] = offer 
}

func (mockDB) removeOffer(sku string) {
	delete(offers, sku)
}