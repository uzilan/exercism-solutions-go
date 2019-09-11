// Package grains provides a solution to the grains exercise of the Go track in https://exercism.io
package grains

import (
	"errors"
	"math/big"
)

var one = big.NewInt(1)
var two = big.NewInt(2)
var sixtyfour = big.NewInt(64)

// Square calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, errors.New("wrong input")
	}

	var bigInput = big.NewInt(int64(input - 1))

	var result big.Int
	result.Exp(two, bigInput, nil)
	return result.Uint64(), nil
}

// Total calculates the total number of grains of wheat on a chessboard given that the number on each square doubles.
func Total() uint64 {
	var result big.Int
	result.Exp(two, sixtyfour, nil)
	result.Sub(&result, one)
	return result.Uint64()
}
