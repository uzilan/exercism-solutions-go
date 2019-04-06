// Package pangram provides a solution to the Pangram exercise of the Go track in https://exercism.io
package pangram

import (
	"strings"
)

// IsPangram determines if a sentence is a pangram.
func IsPangram(input string) bool {
	lowerInput := strings.ToLower(input)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		if !strings.ContainsRune(lowerInput, r) {
			return false
		}
	}
	return true
}
