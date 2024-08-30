package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* Question: Define a struct which has at least two methods: getString: to get a string from console input printString: to print the string in upper case.
Also please include simple test function to test the class methods. */

type ConsoleReader struct {
	s string
}

func (cs *ConsoleReader) ReadString() {
	buf := bufio.NewReader(os.Stdin)
	s, err := buf.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading: %s\n", err)
		return
	}
	cs.s = strings.TrimSpace(s)
}

func (cs ConsoleReader) PrintString() {
	fmt.Println(cs.s)
}

func main() {
	var cs ConsoleReader
	fmt.Println("Type something: ")
	cs.ReadString()
	fmt.Println("Here's what you typed:")
	cs.PrintString()
}
