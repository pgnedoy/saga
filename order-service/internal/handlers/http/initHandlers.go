package handlers

import (
	"github.com/pgnedoy/saga/order-service/internal/repository"
	"github.com/pgnedoy/saga/order-service/internal/usecases"
)

func InitHandlers() (*Handlers, error) {
	repo := repository.NewRepoAdapter()
	createOrder, _ := usecases.NewCreateOrder(&usecases.CreateOrderConfig{Repo: repo})
	handlers, err := NewHandlers(&HandlersConfig{
		CreateOrder: createOrder,
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
