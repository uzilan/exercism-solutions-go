// Package isbn provides a solution to the ISBN Verifier exercise of the Go track in https://exercism.io
package isbn

import "strings"

// IsValidISBN checks if the provided string is a valid ISBN-10
func IsValidISBN(input string) bool {
	cleaned := strings.Replace(input, "-", "", -1)
	if len(cleaned) != 10 {
		return false
	}

	sum := 0
	for i, r := range cleaned {
		n := parseRune(i, r)
		sum = sum + n*(10-i)
	}
	return sum%11 == 0
}

func parseRune(i int, r rune) int {
	if i == 9 && r == 'X' {
		return 10
	}
	return int(r - '0')
}
