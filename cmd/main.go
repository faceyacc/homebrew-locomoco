/*
Copyright © 2023 NAME HERE justfacey@gmail.com
*/
package cmd

import (
	"locomoco/internals"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loco-moco",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		dotFile := internals.GetShowMeDotFilePath()

		folder, _ := cmd.Flags().GetString("add")
		email, _ := cmd.Flags().GetString("email")
		user, _ := cmd.Flags().GetString("user")

		if folder != "" {
			scan(folder)
			return
		}

		if email == "" || user == "" {

			email, _ = GetUserInfo(dotFile)
		} else if email != "" && user != "" {

			SetUserInfo(email, user)

		}

		stats(email)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var folder string
	var email string
	var user string

	rootCmd.PersistentFlags().StringVar(&folder, "add", "", "Adds a folder to the list to scan.")
	rootCmd.PersistentFlags().StringVar(&email, "email", "", "The email to scan.")
	rootCmd.PersistentFlags().StringVar(&user, "user", "", "Show list of repos for given user")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
