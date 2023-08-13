package services

import (
	"log"

	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/atcheri/tavern-ddd-go/domain/product"
	customerRepo "github.com/atcheri/tavern-ddd-go/infrastructure/db/memory/customer"
	productRepo "github.com/atcheri/tavern-ddd-go/infrastructure/db/memory/product"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customerRepo customer.CustomerRepository
	productRepo  product.ProductRepository
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
	cr := customerRepo.NewCustomerRepo()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := productRepo.NewProductRepo()

		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}

		os.productRepo = pr
		return nil
	}
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := os.customerRepo.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []aggregate.Product
	var total float64
	for _, id := range productsIDs {
		p, err := os.productRepo.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return total, nil
}

func (os *OrderService) AddCustomer(c aggregate.Customer) (uuid.UUID, error) {
	err := os.customerRepo.Add(c)
	if err != nil {
		return uuid.Nil, err
	}
	return c.GetID(), nil
}

func (os *OrderService) AddNewCustomer(name string) (uuid.UUID, error) {
	c, err := aggregate.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	return os.AddCustomer(c)
}
