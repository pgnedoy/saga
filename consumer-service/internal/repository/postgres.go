package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type RepoAdapter struct {}

func NewRepoAdapter() *RepoAdapter {
	return &RepoAdapter{}
}

func (ra *RepoAdapter) FindConsumerByID(ctx context.Context, consumerID string) (*data.Consumer, error) {
	return &data.Consumer{}, nil
}

func (ra *RepoAdapter) SaveConsumer(ctx context.Context, consumer data.Consumer) error {
	return nil
}