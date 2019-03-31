// Package scrabble provides a solution to the scrabble score exercise of the Go track in https://exercism.io
package scrabble

import "unicode"

// Score calculates the scrabble score for a given string
func Score(input string) int {
	result := 0
	for _, letter := range input {
		switch unicode.ToUpper(letter) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			result++
		case 'D', 'G':
			result += 2
		case 'B', 'C', 'M', 'P':
			result += 3
		case 'F', 'H', 'V', 'W', 'Y':
			result += 4
		case 'K':
			result += 5
		case 'J', 'X':
			result += 8
		case 'Q', 'Z':
			result += 10
		}
	}
	return result
}
