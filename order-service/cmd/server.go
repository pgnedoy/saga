package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pgnedoy/saga/core/http"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/order-service/config/postgres"
	httphandlers "github.com/pgnedoy/saga/order-service/internal/handlers/http"
	"github.com/pgnedoy/saga/order-service/internal/repository"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the order-service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			cancel()
		}()

		postgres.RunMigrations()

		conn := postgres.GetConnection(&postgres.ConnConfig{Url: os.Getenv("DB_URL")})
		defer func() {
			conn.Close()
		}()

		repo := repository.NewRepoAdapter(&repository.RepoAdapterConfig{DB: conn})
		handlers, err := httphandlers.InitHandlers(&httphandlers.InitHandlersConfig{
			Repo: repo,
		})
		if err != nil {
			log.Panic(ctx, "init handlers error", log.WithError(err))
		}

		r := mux.NewRouter()
		r.HandleFunc("/health-check", handlers.HealthCheck)
		r.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
		s, err := http.NewServer(r)

		if err != nil {
			log.Error(ctx, "error creating server", log.WithError(err))
		}

		portStr, passed := os.LookupEnv("APP_PORT")
		if passed == false {
			log.Panic(ctx, "APP_PORT is required")
		}
		port, _ := strconv.Atoi(portStr)
		log.Info(ctx, fmt.Sprintf("order-server has beed started on port %d", port))
		err = s.Run(context.Background(), port)
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
