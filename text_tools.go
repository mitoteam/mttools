// Package app contains main application functionality.
package mttools

import (
	"strings"

	lorem "github.com/drhodes/golorem"
)

func RandomWord(minLength, maxLength int, capitalize bool) string {
	word := lorem.Word(minLength, maxLength)

	if capitalize {
		word = strings.Title(word)
	}

	return word
}

func RandomWords(count, minLength, maxLength int, capitalize bool) string {
	if count < 1 {
		panic("Count should be larger than 0")
	}

	result := ""

	for count > 0 {
		count--

		word := lorem.Word(minLength, maxLength)

		if capitalize {
			word = strings.Title(word)
		}

		result = result + word

		if count > 0 {
			result = result + " "
		}
	}

	return result
}
