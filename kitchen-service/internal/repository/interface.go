package repository

import (
	"context"

	"github.com/pgnedoy/saga/core/data"
)

type Repository interface {
	FindTicketByID(ctx context.Context, ticketID string) (*data.Ticket, error)
	CreateTicket(ctx context.Context, ticket data.Ticket) error
	UpdateTicket(ctx context.Context, ticket data.Ticket) error
}
