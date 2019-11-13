// Package prime provides a solution to the Prime Factors exercise of the Go track in https://exercism.io
package prime

// Factors compute the prime factors of a given natural number
func Factors(input int64) []int64 {
	var factors []int64
	i := int64(2)
	max := input
	for i <= max {
		if max%i == 0 {
			factors = append(factors, i)
			max /= i
			i = 2
			continue
		}
		i++
	}
	return factors
}
