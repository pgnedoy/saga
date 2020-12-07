package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

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

func (r *RepoAdapter) FindOrderByID(ctx context.Context, orderID string) (*data.Order, error) {
	var order data.Order

	err := r.db.Get(&order, "SELECT * FROM orders WHERE id=$1", orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Error(ctx, "error fetching order by id", log.WithError(err))
		return nil, err
	}

	return &order, nil
}

func (r *RepoAdapter) CreateOrder(ctx context.Context, order data.Order) error {
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
		return err
	}

	return nil
}

func (r *RepoAdapter) UpdateOrder(ctx context.Context, order data.Order) error {
	updateOrder, err := r.db.Prepare("UPDATE orders SET name=$1, quantity=$2, status=$3, updated_at=$4 consumer_id=$5 WHERE id=$6")
	tx, err := r.db.Begin()
	if err != nil {
		log.Error(ctx, "error transaction begging", log.WithError(err))
		return err
	}

	_, err = tx.StmtContext(ctx, updateOrder).Exec(order.Name, order.Quantity, order.Status, time.Now(), order.ConsumerID, order.ID)

	if errors.Is(err, sql.ErrNoRows) {
		log.Error(ctx, "error updating order", log.WithError(err))
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Error(ctx, "error transaction committing", log.WithError(err))
		return err
	}

	return nil
}