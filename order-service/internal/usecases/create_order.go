package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/order-service/internal/repository"
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

func (co *CreateOrder) Execute(ctx context.Context, order data.Order) error {
	order.Status = data.StatusPending
	err := co.repo.CreateOrder(ctx, order)
	if err != nil {
		log.Info(ctx, "error order creating", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
