package services

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNegativeBill = errors.New("a bill cannot have a negative value")
	ErrZeroBill     = errors.New("a bill cannot have a value of zero")
	ErrTooHighBill  = errors.New("a bill cannot be greater than the max")
	MaxTotalBill    = 10000
)

type BillSender interface {
	Send(uuid.UUID, float64) error
}

type BillingService struct {
	orderService *OrderService
	sender       BillSender
}

func NewBillingService(os *OrderService, s BillSender) (*BillingService, error) {
	return &BillingService{
		orderService: os,
		sender:       s,
	}, nil

}

func (bs *BillingService) BillCustomer(customer uuid.UUID, total float64) error {
	if total < 0 {
		return ErrNegativeBill
	}
	if total == 0 {
		return ErrZeroBill
	}

	if total > float64(MaxTotalBill) {
		return ErrTooHighBill
	}

	err := bs.sender.Send(customer, total)
	if err != nil {
		return err
	}

	return nil
}
