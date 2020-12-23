package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pgnedoy/saga/core/log"
)

type ConnConfig struct {
	Url string
}


func GetConnection(cfg *ConnConfig) *sqlx.DB {
	ctx := context.Background()

	if cfg == nil {
		log.Panic(ctx, "connection config required")
	}

	if len(cfg.Url) == 0 {
		log.Panic(ctx, "connection url is empty")
	}

	connConfig, err := pgx.ParseConfig(cfg.Url)
	if err != nil {
		log.Panic(ctx, "error parsing db config", log.WithError(err))
	}

	nativeDB := stdlib.OpenDB(*connConfig)
	conn := sqlx.NewDb(nativeDB, "pgx")

	if err = conn.Ping(); err != nil {
		log.Panic(ctx, "error connection to db", log.WithError(err))
	}

	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(5)

	return conn
}
