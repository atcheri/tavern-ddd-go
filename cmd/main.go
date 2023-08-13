// Package main runs the tavern and performs an Order
package main

import (
	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/atcheri/tavern-ddd-go/infrastructure/sender"
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

	// Create tavern service
	tavern, err := services.NewTavernService(
		services.WithOrderService(os),
		services.WithBillingService(os, &sender.LogSender{}),
	)
	if err != nil {
		panic(err)
	}

	uid, err := os.AddNewCustomer("Mister who")
	if err != nil {
		panic(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// Execute Order and send the bill to the customer
	err = tavern.Order(uid, order)
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
