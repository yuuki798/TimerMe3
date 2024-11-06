package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yuuki798/TimerMe3/cmd/server"
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
