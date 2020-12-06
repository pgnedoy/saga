package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/consumer-service/internal/repository"
	"github.com/pgnedoy/saga/core/log"
)

type VerifyConsumer struct {
	repo repository.Repository
}

type VerifyConsumerConfig struct {
	Repo repository.Repository
}

func NewVerifyConsumer(cfg *VerifyConsumerConfig) (*VerifyConsumer, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &VerifyConsumer{
		repo: cfg.Repo,
	}, nil
}

func (co *VerifyConsumer) Execute(ctx context.Context) {
	log.Info(ctx, "verify consumer usecase")
}
