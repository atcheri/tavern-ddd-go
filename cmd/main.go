// Package main runs the tavern and performs an Order
package main

import (
	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/atcheri/tavern-ddd-go/services"
	"github.com/google/uuid"
)

func main() {
	products := productInventory()

	// Create Order Service to use in tavern
	os, err := services.NewOrderService(
		services.WithMemoryCustomerRepository(),
		services.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	// TODO: Create the Billing Service to use in the tavern
	// bs, err := services.NewBillingService(...)

	// Create tavern service
	tavern, err := services.NewTavernService(
		services.WithOrderService(os),
		services.WithBillingService(os),
	)
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Mister who")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}

	// Bill the customer
	err = tavern.Bill(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := aggregate.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := aggregate.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []aggregate.Product{
		beer, peenuts, wine,
	}
	return products
}
