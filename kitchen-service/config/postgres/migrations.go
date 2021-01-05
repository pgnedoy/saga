package postgres

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pgnedoy/saga/core/log"

	"os"
)

func RunMigrations() {
	dir, _ := os.Getwd()

	sourceUrl := "file:///" + dir + "/db/migrations"

	m, err := migrate.New(
		sourceUrl,
		os.Getenv("DB_URL"),
	)
	if err != nil {
		log.Fatal(context.Background(), "error running migrations", log.WithError(err))
	}
	err = m.Up()
	if  err != nil && err.Error() != "no change" {
		log.Fatal(context.Background(), "error running migrations", log.WithError(err))
	}
}