package main

import (
	"fmt"
	"strings"
)

// isValidSingleLineComment checks if the string is a valid single-line comment
func isValidSingleLineComment(comment string) bool {
	return strings.HasPrefix(comment, "//")
}

// isValidMultiLineComment checks if the string is a valid multi-line comment
func isValidMultiLineComment(comment string) bool {
	return strings.HasPrefix(comment, "/*") && strings.HasSuffix(comment, "*/")
}

// isValidCommentSection checks if the string is within valid comment sections
func isValidCommentSection(comment string) bool {
	comment = strings.TrimSpace(comment)
	return isValidSingleLineComment(comment) || isValidMultiLineComment(comment)
}

func main() {
	// Test cases
	comments := []string{
		"// This is a single-line comment",
		"/* This is a multi-line comment */",
		"/* This is an invalid multi-line comment",
		"Some random text",
		"// Another single-line comment",
		"/* Another multi-line comment */",
		"/* Unclosed multi-line comment",
	}

	for _, comment := range comments {
		if isValidCommentSection(comment) {
			fmt.Printf("'%s' is within a valid comment section.\n", comment)
		} else {
			fmt.Printf("'%s' is not within a valid comment section.\n", comment)
		}
	}
}
