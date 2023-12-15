/*
Copyright © 2023 NAME HERE justfacey@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var showmeCmd = &cobra.Command{
	Use:   "showme",
	Short: "Shows you a list of your repos...",
	Long:  `Shows you a list of your repos with data (i.e. Name, Description, Created At)`,
	Run: func(cmd *cobra.Command, args []string) {
		newUser, _ := cmd.Flags().GetString("newUser")
		user, _ := cmd.Flags().GetString("user")

		if user == "" && newUser != "" {
			user = newUser

		} else if user == "" {
			_, user = GetUserInfo()
		}

		data := ShowMeRepos(user)
		printData(data)
	},
}

func init() {
	rootCmd.AddCommand(showmeCmd)

	var newUser string

	showmeCmd.PersistentFlags().StringVar(&newUser, "newUser", "", "Show list of repos for given user")

}
