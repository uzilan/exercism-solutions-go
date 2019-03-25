// Package romannumerals provides a solution to the roman numerals exercise of the Go track in https://exercism.io
package romannumerals

import (
	"errors"
	"strings"
)

// ToRomanNumeral translate a given number in arabic notation into roman numerals
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 {
		return "", errors.New("input is equal or less than zero")
	}
	if arabic > 3000 {
		return "", errors.New("input is higher than 3000")
	}

	result := ""
	switch {
	case arabic > 1000:
		s, _ := ToRomanNumeral(arabic - 1000)
		result = "M" + s
	case arabic > 100:
		s, _ := ToRomanNumeral(arabic % 100)
		result = roman(arabic/100, "C", "D", "M") + s
	case arabic > 10:
		s, _ := ToRomanNumeral(arabic % 10)
		result = roman(arabic/10, "X", "L", "C") + s
	case arabic >= 1:
		result = roman(arabic, "I", "V", "X")
	}
	return result, nil
}

func roman(number int, one string, five string, ten string) string {
	switch number {
	case 10:
		return ten
	case 9:
		return one + ten
	case 6, 7, 8:
		return five + strings.Repeat(one, number-5)
	case 5:
		return five
	case 4:
		return one + five
	case 1, 2, 3:
		return strings.Repeat(one, number)
	}
	return ""
}
