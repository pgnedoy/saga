package cmd

import (
	"context"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/pgnedoy/saga/core/http"
	"github.com/pgnedoy/saga/core/log"
	"github.com/pgnedoy/saga/kitchen-service/internal/handlers"
	"github.com/spf13/cobra"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the kitchen-service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer func() {
			cancel()
		}()

		handlers, err := handlers.InitHandlers()
		if err != nil {
			log.Panic(ctx, "init handlers error", log.WithError(err))
		}

		r := mux.NewRouter()
		r.HandleFunc("/health-check", handlers.HealthCheck)
		s, err := http.NewServer(r)

		if err != nil {
			log.Error(ctx, "error creating server", log.WithError(err))
		}

		port := 5001
		log.Info(ctx, fmt.Sprintf("order-server has beed started on port %d", port))
		err = s.Run(context.Background(), port)
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
}

