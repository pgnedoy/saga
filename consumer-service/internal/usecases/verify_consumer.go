package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/consumer-service/internal/repository"
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

func (vc *VerifyConsumer) Execute(ctx context.Context, consumerID string) (bool, error) {
	consumer, err := vc.repo.FindConsumerByID(ctx, consumerID)
	if err != nil {
		return false, err
	}
	if len(consumer.FirstName) != 0 && len(consumer.SecondName) != 0 && len(consumer.Phone) != 0 {
		return true, nil
	}
	return false, nil
}
