// Package luhn provides a solution to the luhn exercise of the Go track in https://exercism.io
package luhn

import (
	"strings"
)

// Valid returns true it given number is valid per the Luhn formula or false otherwise.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) < 2 {
		return false
	}

	sum := 0
	double := len(input)%2 == 0

	for _, ru := range input {
		runeValue := int(ru - '0')
		if runeValue < 0 || runeValue > 9 {
			return false
		}

		if double {
			runeValue *= 2
			if runeValue > 9 {
				runeValue -= 9
			}
		}
		sum += runeValue
		double = !double
	}

	return sum%10 == 0
}
