// Package lsproduct provides a solution to the Largest Series Product exercise
// of the Go track in https://exercism.io
package lsproduct

import (
	"errors"
	"regexp"
)

// LargestSeriesProduct calculates the largest product for a contiguous substring of
// digits of length n, given a string of digits,
func LargestSeriesProduct(digits string, span int) (int, error) {
	if len(digits) < span {
		return -1, errors.New("span must be smaller than string length")
	}

	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	match, _ := regexp.MatchString(".*[a-zA-Z].*", digits)
	if match {
		return -1, errors.New("digits input must only contain digits")
	}

	highest := 0
	for i := 0; i <= len(digits)-span; i++ {
		subString := digits[i : i+span]
		sum := 1
		for _, r := range subString {
			sum *= int(r - '0')
		}
		if sum > highest {
			highest = sum
		}
	}
	return highest, nil
}
