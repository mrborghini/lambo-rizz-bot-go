package main

import (
	"fmt"
	"lambo-rizz-bot-go/api"
)

func main() {
	log := api.NewLogger("main")

	config := api.GetConfig()

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
