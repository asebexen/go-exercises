package main

import (
	"fmt"
	"os"
	"strconv"
)

/* Question: With a given integral number n, write a program to generate a dictionary that contains (i, i*i) such that i is an integral number between 1 and n (both included).
and then the program should print the dictionary. Suppose the following input is supplied to the program: 8
Then, the output should be: {1: 1, 2: 4, 3: 9, 4: 16, 5: 25, 6: 36, 7: 49, 8: 64} */

func genMap(n int) map[int]int {
	if n < 1 {
		return make(map[int]int, 0)
	}

	result := make(map[int]int)
	for i := 1; i <= n; i++ {
		result[i] = i * i
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a number.")
		return
	}

	arg := os.Args[1]
	n, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("Failed to convert argument %s to int.\n", arg)
		return
	}

	result := genMap(n)
	fmt.Printf("%v\n", result)
}
