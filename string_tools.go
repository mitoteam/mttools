package mttools

import (
	"regexp"
	"strings"
)

// Splits string into array of strings.
// Each argument in string is single word or number, or string in double quotes.
func SplitArgumentsString(str string) []string {
	str = strings.TrimSpace(str)

	// \p{L} = any unicode letter
	re := regexp.MustCompile(`([\p{L}\d_\-\+\@\|\!\/]+)|(".*?[^\\]")|("")`)

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
