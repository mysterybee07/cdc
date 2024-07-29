package main

import (
	"fmt"
)

// Operator types
const (
	ADD = "+"
	SUB = "-"
	MUL = "*"
	DIV = "/"
)

// Quadruple represents a basic intermediate code
type Quadruple struct {
	Operator string
	Arg1     string
	Arg2     string
	Result   string
}

// GenerateIntermediateCode generates intermediate code for an arithmetic expression
func GenerateIntermediateCode(expression string) []Quadruple {
	// Dummy implementation for the sake of example
	// Parsing and creating quadruples would be more complex in a real scenario
	return []Quadruple{
		{Operator: ADD, Arg1: "a", Arg2: "b", Result: "t1"},
		{Operator: MUL, Arg1: "t1", Arg2: "c", Result: "t2"},
	}
}

func main() {
	expression := "a + b * c"
	code := GenerateIntermediateCode(expression)

	fmt.Println("Intermediate Code:")
	for _, quad := range code {
		fmt.Printf("(%s, %s, %s, %s)\n", quad.Operator, quad.Arg1, quad.Arg2, quad.Result)
	}
}
