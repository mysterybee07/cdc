// WAP to implement FOLLOW of grammar.
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

func computeFollow(grammar Grammar, startSymbol string) map[string]map[string]bool {
	first := computeFirst(grammar)
	follow := make(map[string]map[string]bool)
	follow[startSymbol] = map[string]bool{"$": true}

	for {
		updated := false
		for nonTerminal, productions := range grammar {
			for _, production := range productions {
				trailer := make(map[string]bool)
				for k := range follow[nonTerminal] {
					trailer[k] = true
				}
				for i := len(production) - 1; i >= 0; i-- {
					symbol := string(production[i])
					if _, ok := grammar[symbol]; ok {
						if follow[symbol] == nil {
							follow[symbol] = make(map[string]bool)
						}
						for k := range trailer {
							if !follow[symbol][k] {
								follow[symbol][k] = true
								updated = true
							}
						}
						if _, ok := first[symbol]["ε"]; ok {
							for k := range first[symbol] {
								if k != "ε" {
									trailer[k] = true
								}
							}
						} else {
							trailer = make(map[string]bool)
							for k := range first[symbol] {
								trailer[k] = true
							}
						}
					} else {
						trailer = make(map[string]bool)
						trailer[symbol] = true
					}
				}
			}
		}
		if !updated {
			break
		}
	}
	return follow
}

func main() {
	grammar := Grammar{
		"S": {"AB"},
		"A": {"aA", "ε"},
		"B": {"bB", "ε"},
	}

	startSymbol := "S"
	followSets := computeFollow(grammar, startSymbol)
	for nonTerminal, followSet := range followSets {
		fmt.Printf("FOLLOW(%s) = { ", nonTerminal)
		for symbol := range followSet {
			fmt.Printf("%s ", symbol)
		}
		fmt.Println("}")
	}
}
