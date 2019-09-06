// Package luhn provides a solution to the luhn exercise of the Go track in https://exercism.io
package luhn

import (
	"regexp"
	"strings"
)

// Valid returns true it given number is valid per the Luhn formula or false otherwise.
func Valid(input string) bool {
	noSpaces := strings.ReplaceAll(input, " ", "")
	re := regexp.MustCompile(`[^0-9]`)

	if len(noSpaces) < 2 || re.MatchString(noSpaces) {
		return false
	}

	reversed := reverse(noSpaces)
	sum := int32(0)

	for index, ru := range reversed {
		sum += findRuneValue(index, ru)
	}

	return sum%10 == 0
}

func reverse(input string) string {
	reversed := ""
	for _, ru := range input {
		reversed = string(ru) + reversed
	}
	return reversed
}

func findRuneValue(index int, ru rune) int32 {
	runeValue := ru - '0'
	if (index+1)%2 != 0 {
		return runeValue
	}

	multiplied := runeValue * 2
	if multiplied > 9 {
		multiplied -= 9
	}
	return multiplied
}
