package http

import (
	"context"

	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/repository"
	"github.com/pgnedoy/saga/kitchen-service/internal/usecases"
)

type InitHandlersConfig struct {
	Repo repository.Repository
}

func InitHandlers(cfg *InitHandlersConfig) (*Handlers, error) {
	if cfg == nil {
		log.Panic(context.Background(), "http handlers config is required")
	}

	if cfg.Repo == nil {
		log.Panic(context.Background(), "repo is required for http handlers config")
	}
	
	createTicket, _ := usecases.NewCreateTicket(&usecases.CreateTicketConfig{Repo: cfg.Repo})
	approveTicket, _ := usecases.NewApproveTicket(&usecases.ApproveTicketConfig{Repo: cfg.Repo})
	rejectTicket, _ := usecases.NewRejectTicket(&usecases.RejectTicketConfig{Repo: cfg.Repo})
	handlers, err := NewHandlers(&HandlersConfig{
		CreateTicket:  createTicket,
		ApproveTicket: approveTicket,
		RejectTicket:  rejectTicket,
	})

	if err != nil {
		return nil, err
	}

	return handlers, nil
}
