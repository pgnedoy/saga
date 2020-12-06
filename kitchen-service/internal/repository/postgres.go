package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type RepoAdapter struct {}

func NewRepoAdapter() *RepoAdapter {
	return &RepoAdapter{}
}

func (ra *RepoAdapter) FindTicketByID(ctx context.Context, ticketID string) (*data.Ticket, error) {
	return &data.Ticket{}, nil
}

func (ra *RepoAdapter) SaveTicket(ctx context.Context, ticket data.Ticket) error {
	return nil
}