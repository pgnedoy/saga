package usecases

import (
	"context"

	"github.com/pgnedoy/saga/core/log"
)

type CreateOrder struct {}

func NewCreateOrder() *CreateOrder {
	return &CreateOrder{}
}

func (co *CreateOrder) Execute(ctx context.Context) {
	log.Info(ctx, "create order usecase")
}
