package memory

import (
	"fmt"
	"sync"

	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]customer.Customer
	mutex     sync.Mutex
}

func NewCustomerRepo() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c customer.Customer) error {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	if mr.customers == nil {
		mr.customers = make(map[uuid.UUID]customer.Customer)
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	mr.customers[c.GetID()] = c

	return nil
}

func (mr *MemoryCustomerRepository) Update(c customer.Customer) error {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.customers[c.GetID()] = c

	return nil
}
