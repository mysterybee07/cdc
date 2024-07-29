package main

import (
	"fmt"
)

const stackSize = 30

var stack [stackSize]rune
var top int = -1

// Push character onto stack
func push(c rune) {
	if top < stackSize-1 {
		top++
		stack[top] = c
	} else {
		fmt.Println("Stack overflow")
		panic("Stack overflow")
	}
}

// Pop character from stack
func pop() rune {
	if top != -1 {
		c := stack[top]
		top--
		return c
	}
	return 'x' // Return 'x' if stack is empty
}

// Print current stack status
func printStat() {
	fmt.Print("\n$")
	for i := 0; i <= top; i++ {
		fmt.Print(string(stack[i]))
	}
}

// Main function
func main() {
	s1 := "id+id" // Predefined expression
	l := len(s1)

	fmt.Println("LR PARSING")
	fmt.Printf("Expression: %s", s1)
	fmt.Print("\n$")

	// Process input and simulate stack operations
	var i int
	var token rune
	for i = 0; i < l; i++ {
		token = rune(s1[i])
		if token == 'i' && i+1 < l && s1[i+1] == 'd' {
			fmt.Print("E")
			push('E')
			printStat()
			i++ // Skip next character
		} else if token == '+' || token == '*' || token == '/' {
			push(token)
			printStat()
		}
	}

	// Final stack status
	printStat()

	// Reduction process
	for top != -1 {
		ch1 := pop()
		if ch1 == 'x' {
			fmt.Print("\n$")
			break
		}
		if ch1 == '+' || ch1 == '/' || ch1 == '*' {
			ch3 := pop()
			if ch3 != 'E' {
				fmt.Println("error")
				panic("Unexpected character during reduction")
			} else {
				push('E')
				printStat()
			}
		}
	}

	// Ensure a newline after the final E
	if top >= 0 && stack[top] == 'E' {
		fmt.Println()
	} else {
		fmt.Println() // Ensure a newline if stack is empty or contains other items
	}
}
