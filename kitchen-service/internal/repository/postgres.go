package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/pgnedoy/saga/core/data"
	"github.com/pgnedoy/saga/core/log"
)


type RepoAdapter struct {
	db *sqlx.DB
}

type RepoAdapterConfig struct {
	DB *sqlx.DB
}

func NewRepoAdapter(cfg *RepoAdapterConfig) *RepoAdapter {
	if cfg == nil {
		log.Panic(context.Background(), "RepoAdapterConfig is required")
	}

	if cfg.DB == nil {
		log.Panic(context.Background(), "RepoAdapterConfig: connection is required")
	}

	return &RepoAdapter{
		db: cfg.DB,
	}
}

func (r *RepoAdapter) FindTicketByID(ctx context.Context, ticketID string) (*data.Ticket, error) {
	var ticket data.Ticket

	err := r.db.Get(&ticket, "SELECT * FROM tickets WHERE id=$1", ticketID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Error(ctx, "error fetching ticket by id", log.WithError(err))
		return nil, err
	}

	return &ticket, nil
}

func (r *RepoAdapter) CreateTicket(ctx context.Context, ticket data.Ticket) error {
	createOrder, err := r.db.Prepare("INSERT INTO orders (name, quantity, status, consumer_id) VALUES ($1, $2, $3, $4)")
	tx, err := r.db.Begin()
	if err != nil {
		log.Error(ctx, "error transaction begging", log.WithError(err))
		return err
	}

	_, err = tx.StmtContext(ctx, createOrder).Exec(order.Name, order.Quantity, order.Status, order.ConsumerID)

	if errors.Is(err, sql.ErrNoRows) {
		log.Error(ctx, "error creating order", log.WithError(err))
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error(ctx, "error transaction committing", log.WithError(err))
		tx.Rollback()
		return err
	}

	return nil
}