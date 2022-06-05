/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"go-embedded-svelte/db"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource",
	Long:  `Create a resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	addCreateUserCommand()
}

func addCreateUserCommand() {

	var username string
	var password string

	// createUserCmd represents the createUser command
	createUserCmd := &cobra.Command{
		Use:   "user",
		Short: "A brief description of your command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			
			db.CreateUser(db.User{Username: username, Password: password})
		},
	}

	createUserCmd.Flags().StringVarP(&username, "username", "u", "", "username to create")
	createUserCmd.Flags().StringVarP(&password, "password", "p", "", "password for user")
	createUserCmd.MarkFlagRequired("username")
	createUserCmd.MarkFlagRequired("password")

	createCmd.AddCommand(createUserCmd)
}
