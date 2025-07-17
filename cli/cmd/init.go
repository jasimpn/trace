package cmd

import (
	"fmt"
	// "os"
	"github.com/spf13/cobra"
	"github.com/jasimpn/trace/core"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new TraceOps repo",
	Run: func(cmd *cobra.Command, args []string) {
		err := core.InitRepo()
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}