package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ipSym    string
	stack    string
	ipPtr    int
	stPtr    int
	lenInput int
	act      string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n GRAMMAR")
	fmt.Println("\n E->E+E\n E->E/E \n E->E*E\n E->a/b")
	fmt.Print("\n Enter the input symbol:")
	ipSym, _ = reader.ReadString('\n')
	ipSym = strings.TrimSpace(ipSym)
	lenInput = len(ipSym)

	fmt.Println("\n\t Stack table")
	fmt.Println("\n Stack\t\t Input symbol\t\t Action")
	fmt.Println("______\t\t ____________\t\t ______\n")
	fmt.Printf(" $\t\t%s$\t\t\t--\n", ipSym)

	act = "shift " + string(ipSym[ipPtr])

	for ipPtr < lenInput {
		stack += string(ipSym[ipPtr])
		ipPtr++
		fmt.Printf(" $%s\t\t%s$\t\t\t%s\n", stack, ipSym[ipPtr:], act)
		if ipPtr < lenInput {
			act = "shift " + string(ipSym[ipPtr])
		}
		check()
		stPtr = len(stack) - 1
	}
	check()
}

func check() {
	flag := false

	// Check for reductions E->a or E->b
	if len(stack) > 0 {
		temp2 := string(stack[len(stack)-1])
		if temp2 == "a" || temp2 == "b" {
			stack = stack[:len(stack)-1] + "E"
			fmt.Printf(" $%s\t\t%s$\t\t\tE->%s\n", stack, ipSym[ipPtr:], temp2)
			flag = true
		}
	}

	// Check for reductions E->E+E, E->E/E, E->E*E
	if len(stack) >= 3 {
		temp2 := stack[len(stack)-3:]
		if temp2 == "E+E" || temp2 == "E/E" || temp2 == "E*E" {
			stack = stack[:len(stack)-3] + "E"
			fmt.Printf(" $%s\t\t%s$\t\t\tE->%s\n", stack, ipSym[ipPtr:], temp2)
			flag = true
		}
	}

	// Check if the final reduction to E is done
	if stack == "E" && ipPtr == lenInput {
		fmt.Printf(" $%s\t\t%s$\t\t\tACCEPT\n", stack, ipSym[ipPtr:])
		os.Exit(0)
	}

	if !flag && ipPtr == lenInput {
		fmt.Printf("%s\t\t\t%s\t\t reject\n", stack, ipSym[ipPtr:])
		os.Exit(0)
	}
}
