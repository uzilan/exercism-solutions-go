// Package acronym provides a solution to the acronym exercise of the Go track in https://exercism.io
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	noSingleQuotes := replace("'", s, "")
	noNonLetters := replace("[-,]", noSingleQuotes, " ")
	noDoubleSpaces := replace("\\W+", noNonLetters, " ")

	split := strings.Split(noDoubleSpaces, " ")
	acronym := ""
	for _, word := range split {
		acronym += string(word[0])
	}
	return strings.ToUpper(acronym)
}

func replace(regex string, s string, replaceWith string) string {
	re, _ := regexp.Compile(regex)
	return re.ReplaceAllString(s, replaceWith)
}
