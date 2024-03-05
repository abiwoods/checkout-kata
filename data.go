package main

import (
	"fmt"
)

type dataSource interface {
	getProduct(string) (Product, error)
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

func (mockDB) getProduct(SKU string) (Product, error) {
	result, ok := products[SKU]

	if !ok {
		return Product{}, fmt.Errorf("product %s not found", SKU)
	}

	return result, nil
}

func (mockDB) getOffer(SKU string) (Offer, error) {
	result, ok := offers[SKU]

	if !ok {
		return Offer{}, fmt.Errorf("offers for product %s not found", SKU)
	}

	return result, nil
}