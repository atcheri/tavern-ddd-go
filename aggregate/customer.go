package aggregate

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
