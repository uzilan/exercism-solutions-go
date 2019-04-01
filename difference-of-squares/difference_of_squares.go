// Package diffsquares provides a solution to the Difference Of Squares exercise of the Go track in https://exercism.io
// Inspired by Gauss and https://brilliant.org/wiki/sum-of-n-n2-or-n3/
package diffsquares

// SquareOfSum calculates the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares calculates the sum of square of the first N natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference Finds the difference between the square of the sum and the sum of the squares of the first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
