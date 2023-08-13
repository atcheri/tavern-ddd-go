package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *TavernService) error

type TavernService struct {
	OrderService   *OrderService
	BillingService *BillingService
}

func NewTavernService(configs ...TavernConfiguration) (*TavernService, error) {
	t := &TavernService{}

	for _, config := range configs {
		err := config(t)
		if err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *TavernService) error {
		t.OrderService = os
		return nil
	}
}

func (t *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
	total, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("Billing the customer %s of %0.0f TavernCoins", customer, total)

	// TODO: add this method, test and implement
	// t.BillingService.Bill(customer, total)

	return nil
}
