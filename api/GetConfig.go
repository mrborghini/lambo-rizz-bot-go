package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type TwitchData struct {
	Nickname string   `json:"Nickname"`
	OAuth    string   `json:"OAuth"`
	Channels []string `json:"Channels"`
}

func GetConfig() TwitchData {
	log := NewLogger("GetConfig");
	jsonData, err := os.ReadFile("config.json")
	if err != nil {
		log.Error(fmt.Sprintf("Error reading JSON file: %f", err))
		os.Exit(1)
	}
	
	var config TwitchData

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Error(fmt.Sprintf("Error parsing JSON file: %f", err))
		os.Exit(1)
	}
	return config
}
