package http

import (
	"context"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/consumer-service/internal/repository"
	"github.com/pgnedoy/saga/consumer-service/internal/usecases"
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

	createConsumer, _ := usecases.NewCreateConsumer(&usecases.CreateConsumerConfig{Repo: cfg.Repo})
	verifyConsumer, _ := usecases.NewVerifyConsumer(&usecases.VerifyConsumerConfig{Repo: cfg.Repo})
	handlers, err := NewHandlers(&HandlersConfig{
		CreateConsumer: createConsumer,
		VerifyConsumer: verifyConsumer,
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
