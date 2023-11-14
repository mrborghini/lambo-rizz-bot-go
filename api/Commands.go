package api

import (
	"fmt"
	"math/rand"
)

func Rizz(username string) string {
	randomNumber := rand.Intn(101)
	return fmt.Sprintf("%s has %d%% rizz", username, randomNumber)
}
