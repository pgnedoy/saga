package usecases

import (
	"context"
	"errors"

	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
)

type RejectTicket struct {
	repo repository.Repository
}

type RejectTicketConfig struct {
	Repo repository.Repository
}

func NewRejectTicket(cfg *RejectTicketConfig) (*RejectTicket, error) {
	if cfg == nil {
		return nil, errors.New("")
	}

	if cfg.Repo == nil {
		return nil, errors.New("")
	}

	return &RejectTicket{
		repo: cfg.Repo,
	}, nil
}

func (rt *RejectTicket) Execute(ctx context.Context, ticketID string) error {
	ticket, err := rt.repo.FindTicketByID(ctx, ticketID)
	ticket.Status = data.StatusRejected
	err = rt.repo.UpdateTicket(ctx, *ticket)
	if err != nil {
		log.Info(ctx, "error ticket rejecting", log.WithError(err))
		// todo: return reserved error
		return err
	}
	return nil
}
