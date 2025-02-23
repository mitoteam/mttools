package mttools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AskUserChoiceMultiple(prompt string, options_list []string) (choice_list []int, err error) {
	for k, v := range options_list {
		fmt.Printf("%2d: %s\n", k+1, v)
	}

	//Ask user
	fmt.Print("*** " + prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err = scanner.Err()
	if err != nil {
		return []int{}, err
	}

	// Preprocess
	user_input := strings.TrimSpace(scanner.Text())
	user_input = strings.ReplaceAll(user_input, ",", " ")

	number_string_list := strings.Split(user_input, " ")

	if len(user_input) == 0 {
		return []int{}, nil //empty slice
	}

	// Set answers to slice
	choice_list = make([]int, 0, len(number_string_list))

	for _, v := range number_string_list {
		if len(v) < 1 {
			continue
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			return choice_list, fmt.Errorf("wrong input: %s (%s)", v, err.Error())
		}

		if n < 1 || n > len(options_list) {
			continue
		}

		choice_list = append(choice_list, n-1)
	}

	return UniqueSlice(choice_list), nil
}
