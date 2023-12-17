package api

import (
	"strings"
)

// Only get the channel name
func GetChannel(message string) string {
	splitMessage := strings.Fields(message)
	if len(splitMessage) > 2 {
		return splitMessage[2]
	}
	return ""
}

// Only get the username
func GetUsername(message string) string {
	result := "@"
	splitAt := strings.Split(message, "@")
	if len(splitAt) > 1 {
		splitName := strings.Split(splitAt[1], ".")
		if len(splitName) > 0 {
			return result + splitName[0]
		}
	}
	return ""
}

// Only get the message
func GetMessage(message string) string {
	splitMessage := strings.Split(message, ":")
	if len(splitMessage) > 2 {
		return splitMessage[2]
	}
	return ""
}