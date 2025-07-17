package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/jasimpn/trace/config"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add resources (e.g., repo)",
}

var repoCmd = &cobra.Command{
	Use:   "repo [repo name] [URL of the repo]",
	Short: "Add a repo to local config for simple push",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Define a new repo
		newRepo := config.Repo{
			Repository: args[0],
			URL:        args[1],
		}
		// Add the repo
		err := config.AddRepo(newRepo)
		if err != nil {
			log.Fatalf("Failed to add repo: %v", err)
		}

		fmt.Println("Successfully added repo:", newRepo.Repository)
		

	},
}



func init() {
	addCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(addCmd)
}