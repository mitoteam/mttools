package mttools

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AskUserChoiceMultiple(prompt string, options_list []string, unique bool) (choice_list []int, err error) {
	choice_list, err = _askUser(prompt, options_list)

	if unique {
		return UniqueSlice(choice_list), err
	} else {
		return choice_list, err
	}
}

func _askUser(prompt string, options_list []string) (user_input []int, err error) {
	// Print options and prompt
	for k, v := range options_list {
		fmt.Printf("%2d: %s\n", k+1, v)
	}

	fmt.Print("*** " + prompt)

	//Ask user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err = scanner.Err()
	if err != nil {
		return []int{}, err
	}

	// Preprocess
	user_input_string := strings.TrimSpace(scanner.Text())
	user_input_string = strings.ReplaceAll(user_input_string, ",", " ")

	number_string_list := strings.Split(user_input_string, " ")

	if len(user_input_string) == 0 {
		return []int{}, nil //empty slice
	}

	// Set answers to slice
	user_input = make([]int, 0, len(number_string_list))

	for _, v := range number_string_list {
		if len(v) < 1 {
			continue
		}

		n, err := strconv.Atoi(v)
		if err != nil {
			return user_input, fmt.Errorf("wrong input: %s (%s)", v, err.Error())
		}

		if n < 1 || n > len(options_list) {
			continue
		}

		user_input = append(user_input, n-1)
	}

	return user_input, nil
}
