package mttools

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const numbers = "0123456789"
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allSymbols = numbers + letters

// Returns true if file exists and access is not denied.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandomString(length int) string {
	buffer := make([]byte, length)

	for i := range buffer {
		buffer[i] = allSymbols[rand.Int63()%int64(len(allSymbols))]
	}

	return string(buffer)
}
