// Package isogram provides a solution to the isogram exercise of the Go track in https://exercism.io
package isogram

import "strings"

// IsIsogram determines if a word or phrase is an isogram
func IsIsogram(input string) bool {
	letterMap := map[rune]bool{}
	noHyphens := strings.ReplaceAll(input, "-", "")
	noSpaces := strings.ReplaceAll(noHyphens, " ", "")
	upperCases := strings.ToUpper(noSpaces)
	for _, letter := range upperCases {
		if letterMap[letter] {
			return false
		}
		letterMap[letter] = true
	}
	return true
}
