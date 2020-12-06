package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type Repository interface {
	FindConsumerByID(ctx context.Context, consumerID string) (*data.Consumer, error)
	SaveConsumer(ctx context.Context, consumer data.Consumer) error
}
