package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
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

func (co *ApproveOrder) Execute(ctx context.Context) {
	log.Info(ctx, "approve ticket usecase")
}
