package memory

import (
	"fmt"
	"sync"

	"github.com/atcheri/tavern-ddd-go/aggregate"
	"github.com/atcheri/tavern-ddd-go/domain/customer"
	"github.com/google/uuid"
)

type MemoryCustomerRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	mutex     sync.Mutex
}

func NewCustomerRepo() *MemoryCustomerRepository {
	return &MemoryCustomerRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *MemoryCustomerRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryCustomerRepository) Add(c aggregate.Customer) error {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	if mr.customers == nil {
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}

	mr.customers[c.GetID()] = c

	return nil
}

func (mr *MemoryCustomerRepository) Update(c aggregate.Customer) error {
	mr.mutex.Lock()
	defer mr.mutex.Unlock()

	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.customers[c.GetID()] = c

	return nil
}
