// Package collatzconjecture provides a solution to the Collatz conjecture exercise of the Go track in https://exercism.io
package collatzconjecture

import "errors"

// CollatzConjecture follows the Collatz conjecture for a given input
// and counts the number of steps it takes to reach 1
func CollatzConjecture(input int) (int, error) {
	if input <= 0 {
		return 0, errors.New("Input must be larger than 0")
	}
	return helper(input, 0), nil
}

func helper(input int, steps int) int {
	switch {
	case input == 1:
		return steps
	case input%2 == 0:
		return helper(input/2, steps+1)
	default:
		return helper(input*3+1, steps+1)
	}
}
