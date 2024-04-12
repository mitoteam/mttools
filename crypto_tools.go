package mttools

import (
	"math/rand"
)

const numbers = "0123456789"
const lowerLetters = "abcdefghijklmnopqrstuvwxyz"
const upperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allLetters = lowerLetters + upperLetters
const allSymbols = numbers + allLetters

// Returns random string of given length.
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandomString(length int) string {
	buffer := make([]byte, length)

	for i := range buffer {
		buffer[i] = allSymbols[rand.Int63()%int64(len(allSymbols))]
	}

	return string(buffer)
}

// Returns random letter.
func RandomLetter(upperOnly bool) string {
	var list string
	if upperOnly {
		list = upperLetters
	} else {
		list = allLetters
	}

	letter := list[rand.Intn(len(list))]

	return string(letter)
}
