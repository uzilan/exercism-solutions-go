// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
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
