/*
Copyright © 2023 Ty Facey justfacey@gmail.com
*/
package cmd

import (
	"fmt"
	"locomoco/internals"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "locomoco",
	Short: "Quick way to view your GitHub contributions.",
	Long: `locomoco is a quick and easy way to view you recent git
contributions without having to leave you terminal or IDE.

Now, go push some code!
`,

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
			if !internals.DotFileExist() {
				fmt.Fprintf(os.Stderr, "No locomocostats file initalized")
				os.Exit(1)
			}

			// Scan user's commits
			stats(email)

		} else if email != "" && user != "" {
			SetUserInfo(email, user)
		}

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
