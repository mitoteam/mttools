package mttools

import (
	"regexp"
	"strings"

	lorem "github.com/drhodes/golorem"
)

// Generate random word
func RandomWord(minLength, maxLength int, capitalize bool) string {
	word := lorem.Word(minLength, maxLength)

	if capitalize {
		word = strings.Title(word)
	}

	return word
}

// Generate sentence of random words
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

// Splits string into array of strings.
// Each argument in string is single word or number, or string in double quotes.
func SplitArgumentsString(str string) []string {
	str = strings.TrimSpace(str)

	// \p{L} = any unicode letter
	re := regexp.MustCompile(`([\p{L}\d_\-\+\@\|!]+)|(".*?[^\\]")|("")`)

	result := re.FindAllString(str, -1) //-1 = all

	// remove double quotes
	for index, item := range result {
		if len(item) > 0 && item[0] == '"' {
			item = item[1 : len(item)-1]

			result[index] = item
		}
	}

	//DBG
	//fmt.Printf("%+v\n", result)

	return result
}
