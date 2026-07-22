/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	token = os.Getenv("token")
)

// goBotCmd represents the goBot command
var goBotCmd = &cobra.Command{
	Use:   "goBot",
	Aliases: []string{"start"},
	Short: "Telegram Bot",
	Long:  `A longer description of Telegram Bot`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go_bot started, version %s\n", appVersion)

		bot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  token,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check token. %s\n", err)
			return
		}

		bot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Print(m.Message().Payload, m.Text())
			payload:=m.Message().Payload
			
			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello I'm study go bot %s!", appVersion))
			}

			return err
		})

		bot.Start()
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
