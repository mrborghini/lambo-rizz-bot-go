package api

import (
	"encoding/json"
	"fmt"
	"os"
)

// Get the required data from config.json
type TwitchData struct {
	Nickname string   `json:"Nickname"`
	OAuth    string   `json:"OAuth"`
	Channels []string `json:"Channels"`
}

// Parse config.json
func GetConfig() TwitchData {
	log := NewLogger("GetConfig");
	jsonData, err := os.ReadFile("config.json")

	// Check if json file exists
	if err != nil {
		log.Error(fmt.Sprintf("Couldn't open config.json. Did you name it correctly? %f", err))
		os.Exit(1)
	}
	
	var config TwitchData

	err = json.Unmarshal(jsonData, &config)

	// Check if json configuration is set correctly
	if err != nil {
		log.Error(fmt.Sprintf("Could not parse config.json. Did you format it correctly? %f", err))
		os.Exit(1)
	}

	
	return config
}
