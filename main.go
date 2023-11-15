package main

import (
	"fmt"
	"lambo-rizz-bot-go/api"
	"os"
)

func main() {
	log := api.NewLogger("main")

	config := api.GetConfig()

	if len(config.OAuth) != 30 {
		log.Error("Invalid OAuth token")
		os.Exit(1)
	}

	tc := api.NewTwitchChat(config.Nickname, config.OAuth)

	if len(config.Channels) != 0 {
		for _, channel := range config.Channels {
			tc.JoinChat(channel)
			log.Info(fmt.Sprintf("Joined channel %s", channel))
		}
	} else {
		log.Warning("Channel list is empty. Please add some channels")
	}


	tc.ConnectAsync()


	tc.StartListeningForMessages()


}
