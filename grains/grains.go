// Package grains provides a solution to the grains exercise of the Go track in https://exercism.io
package grains

import (
	"errors"
)

// Square calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("wrong input")
	}

	return 1 << uint64(input-1), nil
}

// Total calculates the total number of grains of wheat on a chessboard given that the number on each square doubles.
func Total() uint64 {
	return 1<<uint64(64) - 1
}
