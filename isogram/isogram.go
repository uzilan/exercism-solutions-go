// Package isogram provides a solution to the isogram exercise of the Go track in https://exercism.io
package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(input string) bool {
	letterMap := map[rune]bool{}

	for _, letter := range input {
		if letter == ' ' || letter == '-' {
			continue
		}
		lower := unicode.ToLower(letter)
		if letterMap[lower] {
			return false
		}
		letterMap[lower] = true
	}
	return true
}
