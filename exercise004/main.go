package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Question: Write a program which accepts a sequence of comma-separated numbers from console and generate an array which contains every number.
Suppose the following input is supplied to the program: 34,67,55,33,12,98 Then, the output should be: ['34', '67', '55', '33', '12', '98'] */

func createIntArray(s string) ([]int, error) {
	tokens := strings.Split(s, ",")
	result := make([]int, len(tokens))
	for i, ele := range tokens {
		var err error
		result[i], err = strconv.Atoi(string(strings.TrimSpace(ele)))
		if err != nil {
			err := fmt.Errorf("error parsing arg %s: element must be an integer", strings.TrimSpace(ele))
			return []int{}, err
		}
	}
	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a list of comma-separated numbers")
		return
	}

	arg := os.Args[1]
	result, err := createIntArray(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v\n", result)
}
