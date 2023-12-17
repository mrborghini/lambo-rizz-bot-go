package api

import (
	"fmt"
	"math/rand"
)

// Get username and return a string with a random number between 0 and 100
func Rizz(username string) string {
	randomNumber := rand.Intn(101)
	return fmt.Sprintf("%s has %d%% rizz", username, randomNumber)
}
