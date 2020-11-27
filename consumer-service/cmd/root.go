package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "consumer-service",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("error in rootCmd.Execute", err)
	}
}

