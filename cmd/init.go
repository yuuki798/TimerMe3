package cmd

import (
	"github.com/spf13/cobra"
	"github.com/trancecho/mundo-be-template/cmd/config"
	"github.com/trancecho/mundo-be-template/cmd/server"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "app",
	Short:        "app",
	SilenceUsage: true,
	Long:         `app`,
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(config.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
