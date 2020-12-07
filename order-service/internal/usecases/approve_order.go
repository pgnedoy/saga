package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/order-service/internal/repository"
)

type ApproveOrder struct {
	repo repository.Repository
}

type ApproveOrderConfig struct {
	Repo repository.Repository
}

func NewApproveOrder(cfg *ApproveOrderConfig) (*ApproveOrder, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &ApproveOrder{
		repo: cfg.Repo,
	}, nil
}

func (co *ApproveOrder) Execute(ctx context.Context, order data.Order) error {
	order.Status = data.StatusApproved
	err := co.repo.UpdateOrder(ctx, order)
	if err != nil {
		log.Info(ctx, "error order approving", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
