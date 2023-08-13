package services

import (
	"math/rand"
	"testing"

	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func randomPrice() float64 {
	return rand.Float64() * 100
}

func fakeProduct(t *testing.T) aggregate.Product {
	p, e := aggregate.NewProduct(faker.Name(), faker.Paragraph(), randomPrice())
	if e != nil {
		t.Fatal(e)
	}

	return p
}

func generateTestProducts(t *testing.T) []aggregate.Product {
	products := make([]aggregate.Product, 0)
	for i := 1; i < 5; i++ {
		products = append(products, fakeProduct(t))
	}

	return products
}

func TestOrderService_NewOrderService(t *testing.T) {
	products := generateTestProducts(t)

	os, err := NewOrderService(WithMemoryCustomerRepository(), WithMemoryProductRepository(products))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer(faker.Name())
	if err != nil {
		t.Error(err)
	}

	err = os.customerRepo.Add(cust)
	if err != nil {
		t.Error(err)
	}

	orders := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), orders)
	if err != nil {
		t.Error(err)
	}
}
