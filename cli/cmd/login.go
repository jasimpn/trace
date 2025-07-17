package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jasimpn/trace/config"
)


var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login the username and email",
	Run: func(cmd *cobra.Command, args []string) {

		var username string
		var email string
		
		fmt.Print("Enter the username: ")
		fmt.Scanf("%s", &username)
		
		fmt.Print("Enter the email: ")
		fmt.Scanf("%s", &email)

		err := config.SetUser(username,email)
		if err != nil {
			fmt.Println("Failed to set user:", err)
			return
		}
		fmt.Println("User info updated!")

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}