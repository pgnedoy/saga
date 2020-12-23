package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
)

type ApproveTicket struct {
	repo repository.Repository
}

type ApproveTicketConfig struct {
	Repo repository.Repository
}

func NewApproveTicket(cfg *ApproveTicketConfig) (*ApproveTicket, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &ApproveTicket{
		repo: cfg.Repo,
	}, nil
}

func (at *ApproveTicket) Execute(ctx context.Context, ticketID string) error {
	ticket, err := at.repo.FindTicketByID(ctx, ticketID)
	ticket.Status = data.StatusApproved
	err = at.repo.UpdateTicket(ctx, *ticket)
	if err != nil {
		log.Info(ctx, "error ticket approving", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
