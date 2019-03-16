// Package leap provides a solution to the leap year exercise of the Go track in https://exercism.io
package leap

// IsLeapYear returns true if the given year is a leap year, or false otherwise
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
