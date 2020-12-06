package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/log"
)

type AuthCreditCard struct {
}

type AuthCreditCardConfig struct {
}

func NewAuthCreditCard(cfg *AuthCreditCardConfig) (*AuthCreditCard, error) {
	if cfg == nil {
		return nil, errors.New("")
	}
	
	return &AuthCreditCard{}, nil
}

func (co *AuthCreditCard) Execute(ctx context.Context) {
	log.Info(ctx, "authorise credit card usecase")
}
