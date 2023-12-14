/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showmeCmd represents the showme command
var showmeCmd = &cobra.Command{
	Use:   "showme",
	Short: "Shows you a list of your repos...",
	Long:  `Shows you a list of your repos with data (i.e. Name, Description, Created At)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(ShowMeRepos())
	},
}

func init() {
	rootCmd.AddCommand(showmeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showmeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showmeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
