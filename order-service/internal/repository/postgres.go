package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type RepoAdapter struct {}

func NewRepoAdapter() *RepoAdapter {
	return &RepoAdapter{}
}

func (ra *RepoAdapter) FindOrderByID(ctx context.Context, orderID string) (*data.Order, error) {
	return &data.Order{}, nil
}

func (ra *RepoAdapter) SaveOrder(ctx context.Context, order data.Order) error {
	return nil
}