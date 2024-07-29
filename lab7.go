package main

import (
	"fmt"
)

// Quadruple represents a basic intermediate code
type Quadruple struct {
	Operator string
	Arg1     string
	Arg2     string
	Result   string
}

// GenerateAssembly generates assembly-like code from intermediate code
func GenerateAssembly(code []Quadruple) []string {
	var assemblyCode []string

	for _, quad := range code {
		switch quad.Operator {
		case "+":
			assemblyCode = append(assemblyCode, fmt.Sprintf("ADD %s, %s, %s", quad.Arg1, quad.Arg2, quad.Result))
		case "-":
			assemblyCode = append(assemblyCode, fmt.Sprintf("SUB %s, %s, %s", quad.Arg1, quad.Arg2, quad.Result))
		case "*":
			assemblyCode = append(assemblyCode, fmt.Sprintf("MUL %s, %s, %s", quad.Arg1, quad.Arg2, quad.Result))
		case "/":
			assemblyCode = append(assemblyCode, fmt.Sprintf("DIV %s, %s, %s", quad.Arg1, quad.Arg2, quad.Result))
		default:
			fmt.Printf("Unknown operator: %s\n", quad.Operator)
		}
	}

	return assemblyCode
}

func main() {
	// Intermediate code example
	code := []Quadruple{
		{Operator: "+", Arg1: "a", Arg2: "b", Result: "t1"},
		{Operator: "*", Arg1: "t1", Arg2: "c", Result: "t2"},
	}

	assemblyCode := GenerateAssembly(code)

	fmt.Println("Assembly Code:")
	for _, line := range assemblyCode {
		fmt.Println(line)
	}
}
