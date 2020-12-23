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

func (r *RepoAdapter) FindConsumerByID(ctx context.Context, consumerID string) (*data.Consumer, error) {
	var consumer data.Consumer

	err := r.db.Get(&consumer, "SELECT * FROM consumers WHERE id=$1", consumerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		log.Error(ctx, "error fetching ticket by id", log.WithError(err))
		return nil, err
	}

	return &consumer, nil
}

func (r *RepoAdapter) SaveConsumer(ctx context.Context, consumer data.Consumer) error {
	createConsumer, err := r.db.Prepare("INSERT INTO consumers (first_name, second_name, phone, email) VALUES ($1, $2, $3, $4)")
	tx, err := r.db.Begin()
	if err != nil {
		log.Error(ctx, "error transaction begging", log.WithError(err))
		return err
	}

	_, err = tx.StmtContext(ctx, createConsumer).Exec(consumer.FirstName, consumer.SecondName, consumer.Phone, consumer.Email)

	if errors.Is(err, sql.ErrNoRows) {
		log.Error(ctx, "error creating consumer", log.WithError(err))
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