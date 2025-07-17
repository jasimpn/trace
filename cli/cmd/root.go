package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trace",
	Short: "TraceOps CLI for versioning and audit",
	Long:  `Versioning platform for AI, protocols, and industrial data with Git-like traceability.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}