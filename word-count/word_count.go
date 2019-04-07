// Package wordcount provides a solution to the Word Count exercise of the Go track in https://exercism.io
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency the frequency of the occurrences of each word
type Frequency map[string]int

// WordCount counts the occurrences of each word in that phrase given a phrase.
func WordCount(input string) Frequency {
	frequency := Frequency{}
	noDashses := regexp.MustCompile("\\s'|'\\s").ReplaceAllString(input, " ")
	cleanInput := regexp.MustCompile("[^a-zA-Z0-9']").ReplaceAllString(noDashses, " ")
	noDoubleSpaces := regexp.MustCompile("\\s+").ReplaceAllString(cleanInput, " ")
	trimmed := strings.TrimSpace(noDoubleSpaces)
	lowerCase := strings.ToLower(trimmed)
	words := strings.Split(lowerCase, " ")
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}
