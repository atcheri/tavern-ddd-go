package services

import (
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(ts *TavernService) error

type TavernService struct {
	orderService   *OrderService
	billingService *BillingService
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
		t.orderService = os
		return nil
	}
}

func WithBillingService(os *OrderService, sender BillSender) TavernConfiguration {
	return func(t *TavernService) error {
		bs, err := NewBillingService(os, sender)
		if err != nil {
			return err
		}

		t.billingService = bs
		return nil
	}
}

func (t *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
	total, err := t.orderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("Billing the customer %s of %0.0f TavernCoins", customer, total)

	t.billingService.BillCustomer(customer, total)

	return nil
}
