package main

import (
	"flag"
	"os"
	"log"

	"github.com/nlopes/slack"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	api := slack.New(token)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	
	var (
		statusText string
		statusEmoji string
	)

	flag.StringVar(&statusText, "text", "default", "customText")
	flag.StringVar(&statusEmoji, "emoji", "+1", "customEmoji")
	flag.Parse()
	statusEmoji = ":" + statusEmoji + ":"

	if err := api.SetUserCustomStatus(statusText, statusEmoji); err != nil {
		logger.Fatalf(`SetUserCustomStatus(%q, %q) = %#v`, statusText, statusEmoji, err)
	}
}
