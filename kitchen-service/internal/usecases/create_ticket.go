package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
)

type CreateOrder struct {
	repo repository.Repository
}

type CreateOrderConfig struct {
	Repo repository.Repository
}

func NewCreateOrder(cfg *CreateOrderConfig) (*CreateOrder, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &CreateOrder{
		repo: cfg.Repo,
	}, nil
}

func (co *CreateOrder) Execute(ctx context.Context) {
	log.Info(ctx, "create ticket usecase")
}
