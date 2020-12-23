package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/consumer-service/internal/repository"
	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
)

type CreateConsumer struct {
	repo repository.Repository
}

type CreateConsumerConfig struct {
	Repo repository.Repository
}

func NewCreateConsumer(cfg *CreateConsumerConfig) (*CreateConsumer, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &CreateConsumer{
		repo: cfg.Repo,
	}, nil
}

func (co *CreateConsumer) Execute(ctx context.Context, Consumer data.Consumer) error {
	err := co.repo.SaveConsumer(ctx, Consumer)
	if err != nil {
		log.Info(ctx, "error consumer creating", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
