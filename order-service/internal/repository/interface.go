package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type Repository interface {
	FindOrderByID(ctx context.Context, orderID string) (*data.Order, error)
	SaveOrder(ctx context.Context, order data.Order) error
}