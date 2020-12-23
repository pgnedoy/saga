package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
)

type CreateTicket struct {
	repo repository.Repository
}

type CreateTicketConfig struct {
	Repo repository.Repository
}

func NewCreateTicket(cfg *CreateTicketConfig) (*CreateTicket, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &CreateTicket{
		repo: cfg.Repo,
	}, nil
}

func (co *CreateTicket) Execute(ctx context.Context, ticket data.Ticket) error {
	ticket.Status = data.StatusPending
	err := co.repo.CreateTicket(ctx, ticket)
	if err != nil {
		log.Info(ctx, "error ticket creating", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
