/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go-embedded-svelte/db"
	"os"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a resource",
	Long:  `Show a resource`,
}

func init() {
	rootCmd.AddCommand(showCmd)
	addShowUserCmd()
}

func addShowUserCmd() {
	var username string
	var id uint
	var withPassword bool

	// showUserCmd represents the showUser command
	var showUserCmd = &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			var user *db.User
			if username != "" {
				user = db.FindUserByUsername(username)
			} else if id != 0 {
				user = db.FindUserById(id)
			} else {
				cmd.Help()
				os.Exit(1)
			}

			fmt.Printf("ID: %v\nUser: %s\n", user.ID, user.Username)
			if withPassword {
				fmt.Printf("Password: %s\n", user.Password)
			}
		},
	}

	showCmd.AddCommand(showUserCmd)

	showUserCmd.Flags().StringVarP(&username, "username", "u", "", "username to create")
	showUserCmd.Flags().UintVar(&id, "id", 0, "password for user")
	showUserCmd.Flags().BoolVar(&withPassword, "with-password", false, "show the user's password")

}
