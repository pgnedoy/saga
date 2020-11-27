package handlers

import (
	"github.com/pgnedoy/saga/order-service/internal/usecases"
)

func InitHandlers() (*Handlers, error) {
	createOrder := usecases.NewCreateOrder()
	handlers, err := NewHandlers(&HandlersConfig{
		CreateOrder: createOrder,
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
