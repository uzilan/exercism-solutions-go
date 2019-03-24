// Package scrabble provides a solution to the scrabble score exercise of the Go track in https://exercism.io
package scrabble

import "strings"

var letterScoreMap = map[rune]int{}

func init() {
	var scoreMap = map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}

	for letters, score := range scoreMap {
		for _, letter := range letters {
			letterScoreMap[letter] = score
		}
	}
}

// Score calculates the scrabble score for a given string
func Score(input string) int {
	result := 0
	for _, letter := range strings.ToUpper(input) {
		result += letterScoreMap[letter]
	}
	return result
}
