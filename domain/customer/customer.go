package customer

import (
	"errors"

	"github.com/atcheri/tavern-ddd-go/entity"
	"github.com/atcheri/tavern-ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPersonName = errors.New("a customer must have a name")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPersonName
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	customer := Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}

	return customer, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
