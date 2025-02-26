package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand.NewSource(time.Now().Unix())

	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
