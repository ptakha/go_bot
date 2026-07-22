/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// goBotCmd represents the goBot command
var goBotCmd = &cobra.Command{
	Use:   "goBot",
	Short: "Telegram Bot",
	Long: `A longer description of Telegram Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("goBot called")
	},
}

func init() {
	rootCmd.AddCommand(goBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
