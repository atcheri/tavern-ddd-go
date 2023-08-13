package services

import (
	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/atcheri/tavern-ddd-go/infrastructure/db/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customerRepo customer.CustomerRepository
}

func NewOrderService(configs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customerRepo = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	c, err := os.customerRepo.Get(customerID)
	if err != nil {
		return err
	}

	return nil
}
