/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"go-embedded-svelte/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List various resources",
	Long:  `List various resources`,
}

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "users",
	Short: "List users",
	Long:  `List users in the database`,
	Run: func(cmd *cobra.Command, args []string) {
		PrintTable("Users", *db.GetUsers())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listUsersCmd)
}
