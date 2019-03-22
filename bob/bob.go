// Package bob provides a solution to the bob exercise of the Go track in https://exercism.io
package bob

import (
	"regexp"
	"strings"
)

// Hey returns an answer based on the input and some strange rules
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	switch {
	case isEmpty(trimmed):
		return "Fine. Be that way!"
	case containsLetters(trimmed) && isUpper(trimmed) && isQuestion(trimmed):
		return "Calm down, I know what I'm doing!"
	case isQuestion(trimmed):
		return "Sure."
	case isUpper(trimmed) && containsLetters(trimmed):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isEmpty(remark string) bool {
	return len(remark) == 0
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func isUpper(remark string) bool {
	return strings.ToUpper(remark) == remark
}

func containsLetters(remark string) bool {
	match, _ := regexp.MatchString(".*[a-zA-Z].*", remark)
	return match
}
