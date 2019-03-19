// Package raindrops provides a solution to the raindrops exercise of the Go track in https://exercism.io
package raindrops

import "strconv"

type code struct {
	num  int
	word string
}

// Convert converts a number to a string, the contents of which depend on the number'word factors
func Convert(input int) string {

	var codes = []code{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	var result string
	for _, c := range codes {
		if divisibleBy(input, c.num) {
			result += c.word
		}
	}

	if len(result) == 0 {
		result += strconv.Itoa(input)
	}

	return result
}

func divisibleBy(input int, number int) bool {
	return input%number == 0
}
