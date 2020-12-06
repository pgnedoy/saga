package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
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

func (co *RejectOrder) Execute(ctx context.Context) {
	log.Info(ctx, "reject ticket usecase")
}
