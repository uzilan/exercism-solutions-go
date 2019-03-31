// Package reverse provides a solution to the Reverse String exercise of the Go track in https://exercism.io
package reverse

// String reverses a string
func String(input string) string {
	reversed := ""
	for _, ru := range input {
		reversed = string(ru) + reversed
	}
	return reversed
}
