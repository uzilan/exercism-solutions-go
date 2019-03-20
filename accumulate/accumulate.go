// Package accumulate provides a solution to the accumulate exercise of the Go track in https://exercism.io
package accumulate

// Accumulate converts every element of a given slice of strings using a given conversion function
func Accumulate(input []string, f func(string) string) []string {
	var converted = make([]string, len(input))
	for index, item := range input {
		converted[index] = f(item)
	}
	return converted
}
