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

func TestOrderService_NewOrderService(t *testing.T) {
	products := make([]aggregate.Product, 0)
	p, e := aggregate.NewProduct(faker.Name(), faker.Paragraph(), randomPrice())
	if e != nil {
		t.Fatal(e)
	}

	products = append(products, p)

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
