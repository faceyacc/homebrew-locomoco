/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// showmeCmd represents the showme command
var showmeCmd = &cobra.Command{
	Use:   "showme",
	Short: "Shows you a list of your repos...",
	Long:  `Shows you a list of your repos with data (i.e. Name, Description, Created At)`,
	Run: func(cmd *cobra.Command, args []string) {
		// user := args[0]
		newUser, _ := cmd.Flags().GetString("newUser")
		user, _ := cmd.Flags().GetString("user")

		if user == "" && newUser != "" {
			user = newUser

		} else if user == "" {
			// TODO call utility func to get stored user name and pass it to ShowMeRepos
			user = GetUserName()
		} else if user != "" {
			SetUsername(user)
		}

		data := ShowMeRepos(user)
		printData(data)
	},
}

func init() {
	rootCmd.AddCommand(showmeCmd)

	// Here you will define your flags and configuration settings.

	var newUser string
	var user string
	showmeCmd.PersistentFlags().StringVar(&newUser, "newUser", "", "Show list of repos for given user")
	showmeCmd.PersistentFlags().StringVar(&user, "user", "", "Show list of repos for given user")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
