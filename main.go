package main

import (
	"os"
	"log"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	const (
		statusText = "test"
		statusEmoji = ":100:"
	)
	if err := api.SetUserCustomStatus(statusText, statusEmoji); err != nil {
		logger.Fatalf(`SetUserCustomStatus(%q, %q) = %#v`, statusText, statusEmoji, err)
	}
}
