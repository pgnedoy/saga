package http

import (
	"context"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/order-service/internal/repository"
	"github.com/pgnedoy/saga/order-service/internal/usecases"
)

type InitHandlersConfig struct {
	Repo repository.Repository
}

func InitHandlers(cfg *InitHandlersConfig) (*Handlers, error) {
	if cfg == nil {
		log.Panic(context.Background(), "http handlers config is required")
	}

	if cfg.Repo == nil {
		log.Panic(context.Background(), "repo is required for http handlers config")
	}
	
	createOrder, _ := usecases.NewCreateOrder(&usecases.CreateOrderConfig{Repo: cfg.Repo})
	handlers, err := NewHandlers(&HandlersConfig{
		CreateOrder: createOrder,
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
