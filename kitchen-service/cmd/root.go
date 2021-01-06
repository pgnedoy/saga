package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "kitchen-service",
}

func Execute() {
	if env := os.Getenv("APP_ENV"); env != "local" && len(env) != 0 {
		time.Sleep(time.Second*10)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error in rootCmd.Execute", err)
	}
}

