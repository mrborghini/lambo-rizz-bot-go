package main

import (
	"fmt"
	"lambo-rizz-bot-go/api"
	"os"
)

func main() {
	// Start a new logger
	log := api.NewLogger("main")

	// Get the twitch info from config.json
	config := api.GetConfig()

	// Check if OAuth token is longer than 30 characters
	if len(config.OAuth) != 30 {
		log.Error("Invalid OAuth token")
		os.Exit(1)
	}

	// Add json data to the twitch chat class
	tc := api.NewTwitchChat(config.Nickname, config.OAuth)

	// Check if there are any channels in config.json. If there is none log a warning
	if len(config.Channels) != 0 {
		for _, channel := range config.Channels {
			tc.JoinChat(channel)
			log.Info(fmt.Sprintf("Joined channel %s", channel))
		}
	} else {
		log.Warning("Channel list is empty. Please add some channels")
	}

	// Connect to Twitch using all the credentials
	tc.ConnectAsync()

	// Wait for new messages to come in
	tc.StartListeningForMessages()
}
