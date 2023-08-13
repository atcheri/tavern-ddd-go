package services

import (
	"testing"

	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
)

func TestTavernService_Order(t *testing.T) {
	// arrange
	products := generateTestProducts(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavernService, err := NewTavernService(
		WithOrderService(os),
		// TODO
		// WithBillingService(),
	)
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
	order := []uuid.UUID{
		products[0].GetID(),
	}

	// act
	err = tavernService.Order(cust.GetID(), order)

	// assert
	if err != nil {
		t.Error(err)
	}
}
