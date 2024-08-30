package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Question: Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.
Suppose the following input is supplied to the program: 8 Then, the output should be: 40320 */

func myFunc(n int) (int, error) {
	if n <= 0 {
		return -1, fmt.Errorf("n must be a positive integer")
	}
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Please provide a number.")
		return
	}

	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Failed to convert argument %s to int.\n", arg)
		return
	}

	result, err := myFunc(n)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%d\n", result)
}
