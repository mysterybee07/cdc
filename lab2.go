// WAP to implement FIRST of grammar.
package main

import (
	"fmt"
	"strings"
)

type Grammar map[string][]string

func computeFirst(grammar Grammar) map[string]map[string]bool {
	first := make(map[string]map[string]bool)

	var firstOf func(symbol string) map[string]bool
	firstOf = func(symbol string) map[string]bool {
		if _, ok := grammar[symbol]; !ok {
			return map[string]bool{symbol: true}
		}
		if set, ok := first[symbol]; ok {
			return set
		}

		firstSet := make(map[string]bool)
		for _, production := range grammar[symbol] {
			for _, char := range production {
				charFirst := firstOf(string(char))
				for k := range charFirst {
					if k != "ε" {
						firstSet[k] = true
					}
				}
				if _, exists := charFirst["ε"]; !exists {
					break
				}
			}
			if strings.Contains(production, "ε") {
				firstSet["ε"] = true
			}
		}
		first[symbol] = firstSet
		return firstSet
	}

	for nonTerminal := range grammar {
		firstOf(nonTerminal)
	}
	return first
}

func main() {
	grammar := Grammar{
		"S": {"AB"},
		"A": {"aA", "ε"},
		"B": {"bB", "ε"},
	}

	firstSets := computeFirst(grammar)
	for nonTerminal, firstSet := range firstSets {
		fmt.Printf("FIRST(%s) = { ", nonTerminal)
		for symbol := range firstSet {
			fmt.Printf("%s ", symbol)
		}
		fmt.Println("}")
	}
}
