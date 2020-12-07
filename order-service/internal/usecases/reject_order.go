package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/order-service/internal/repository"
)

type RejectOrder struct {
	repo repository.Repository
}

type RejectOrderConfig struct {
	Repo repository.Repository
}

func NewRejectOrder(cfg *RejectOrderConfig) (*RejectOrder, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &RejectOrder{
		repo: cfg.Repo,
	}, nil
}

func (co *RejectOrder) Execute(ctx context.Context, order data.Order) error {
	order.Status = data.StatusRejected
	err := co.repo.UpdateOrder(ctx, order)
	if err != nil {
		log.Info(ctx, "error order rejecting", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
