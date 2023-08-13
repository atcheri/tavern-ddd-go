package services

import (
	"testing"

	"github.com/atcheri/tavern-ddd-go/domain/customer"
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

	sender := &mockSender{}
	tavernService, err := NewTavernService(
		WithOrderService(os),
		WithBillingService(os, sender),
	)
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer(faker.Name())
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
	sender.On("Send", cust.GetID(), products[0].GetPrice()).Return(nil)

	// act
	err = tavernService.Order(cust.GetID(), orders)

	// assert
	if err != nil {
		t.Error(err)
	}
}
