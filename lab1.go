// WAP to implement Lexical Analyzer to identify tokens.
package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Token types
const (
	NUMBER   = "NUMBER"
	ASSIGN   = "ASSIGN"
	END      = "END"
	ID       = "ID"
	OP       = "OP"
	NEWLINE  = "NEWLINE"
	SKIP     = "SKIP"
	MISMATCH = "MISMATCH"
)

// Token structure
type Token struct {
	Type   string
	Value  string
	Line   int
	Column int
}

// Token specifications
var tokenSpec = []struct {
	Type string
	Re   *regexp.Regexp
}{
	{NUMBER, regexp.MustCompile(`^\d+(\.\d*)?`)},
	{ASSIGN, regexp.MustCompile(`^=`)},
	{END, regexp.MustCompile(`^;`)},
	{ID, regexp.MustCompile(`^[A-Za-z]+`)},
	{OP, regexp.MustCompile(`^[+\-*/]`)},
	{NEWLINE, regexp.MustCompile(`^\n`)},
	{SKIP, regexp.MustCompile(`^[ \t]+`)},
}

// Tokenize function
func tokenize(code string) ([]Token, error) {
	var tokens []Token
	lines := strings.Split(code, "\n")
	lineNum := 1

	for _, line := range lines {
		column := 1
		for len(line) > 0 {
			var match string
			var tokenType string

			for _, spec := range tokenSpec {
				if loc := spec.Re.FindStringIndex(line); loc != nil && loc[0] == 0 {
					match = line[:loc[1]]
					tokenType = spec.Type
					break
				}
			}

			if match == "" {
				match = line[0:1]
				tokenType = MISMATCH
			}

			if tokenType == NEWLINE {
				break
			} else if tokenType != SKIP {
				if tokenType == MISMATCH {
					return nil, fmt.Errorf("unexpected character %q on line %d", match, lineNum)
				}
				tokens = append(tokens, Token{Type: tokenType, Value: match, Line: lineNum, Column: column})
			}

			column += len(match)
			line = line[len(match):]
		}
		lineNum++
	}

	return tokens, nil
}

func main() {
	code := `
x = 10 + 20;
y = x - 5;
`

	tokens, err := tokenize(code)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, token := range tokens {
		fmt.Printf("%+v\n", token)
	}
}
