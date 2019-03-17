// Package hamming provides a solution to the hamming exercise of the Go track in https://exercism.io
package hamming

import "errors"

// Distance calculates the Hamming difference between two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands have different length")
	}
	var count = 0
	for index := range a {
		if a[index] != b[index] {
			count++
		}
	}

	return count, nil
}
