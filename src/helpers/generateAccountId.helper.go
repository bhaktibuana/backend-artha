package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomAccountId() (string, string, string) {
	rand.Seed(time.Now().UnixNano())

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	usernameLength := rand.Intn(16) + 1
	randomUsername := make([]byte, usernameLength)
	for i := 0; i < usernameLength; i++ {
		randomUsername[i] = charset[rand.Intn(len(charset))]
	}
	randomUsernameString := string(randomUsername)

	tagLineLength := rand.Intn(5) + 1
	randomTagLine := make([]byte, tagLineLength)
	for i := 0; i < tagLineLength; i++ {
		randomTagLine[i] = charset[rand.Intn(len(charset))]
	}
	randomTagLineString := string(randomTagLine)

	randomAccountId := fmt.Sprintf("%s#%s", randomUsernameString, randomTagLineString)

	return randomUsernameString, randomTagLineString, randomAccountId
}
