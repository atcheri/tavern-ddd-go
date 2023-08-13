package services

import (
	"testing"

	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/atcheri/tavern-ddd-go/domain/product"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
)

type mockSender struct {
	mock.Mock
}

func (ms *mockSender) Send(id uuid.UUID, total float64) error {
	ms.Called(id, total)
	return nil
}

func TestBillingService_BillCustomer(t *testing.T) {
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
	bs, err := NewBillingService(os, sender)
	if err != nil {
		t.Error(err)
	}

	customer, err := customer.NewCustomer(faker.Name())
	if err != nil {
		t.Error(err)
	}

	os.AddCustomer(customer)

	orderedProducts := lo.Map(products, func(p product.Product, index int) uuid.UUID {
		return p.GetID()
	})
	total, err := os.CreateOrder(customer.GetID(), orderedProducts)
	if err != nil {
		t.Error(err)
	}

	sender.On("Send", customer.GetID(), total).Return(nil)

	// act
	err = bs.BillCustomer(customer.GetID(), total)

	// assert
	if err != nil {
		t.Error(err)
	}
	sender.AssertExpectations(t)
}
