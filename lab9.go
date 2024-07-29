package main

import (
	"fmt"
	"unicode"
)

// isValidIdentifier checks if the provided identifier is valid
func isValidIdentifier(identifier string) bool {
	if len(identifier) == 0 {
		return false
	}

	// Check the first character
	firstChar := rune(identifier[0])
	if !unicode.IsLetter(firstChar) && firstChar != '_' {
		return false
	}

	// Check the remaining characters
	for _, char := range identifier[1:] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_' {
			return false
		}
	}

	// Check if the identifier is a keyword
	keywords := map[string]bool{
		"break": true, "default": true, "func": true, "interface": true,
		"select": true, "case": true, "defer": true, "go": true, "map": true,
		"struct": true, "chan": true, "else": true, "goto": true, "package": true,
		"switch": true, "const": true, "fallthrough": true, "if": true, "range": true,
		"type": true, "continue": true, "for": true, "import": true, "return": true,
		"var": true, "nil": true, "true": true, "false": true,
	}
	if _, exists := keywords[identifier]; exists {
		return false
	}

	return true
}

func main() {
	// Test cases
	identifiers := []string{
		"validIdentifier", "invalid-identifier", "123invalid", "_valid123", "func",
	}

	for _, id := range identifiers {
		if isValidIdentifier(id) {
			fmt.Printf("'%s' is a valid identifier.\n", id)
		} else {
			fmt.Printf("'%s' is not a valid identifier.\n", id)
		}
	}
}
