// Package sieve provides a solution to the sieve exercise of the Go track in https://exercism.io
package sieve

// Sieve uses the Sieve of Eratosthenes to find all the primes from 2 up to a given number
func Sieve(limit int) []int {
	candidates := make([]bool, limit+1)
	for current := 0; current <= limit; current++ {
		candidates[current] = current >= 2
	}

	for index, value := range candidates {
		if !value {
			continue
		}
		for current := index + 1; current <= limit; current++ {
			if current%index == 0 {
				candidates[current] = false
			}
		}
	}

	var primes []int
	for index, value := range candidates {
		if value {
			primes = append(primes, index)
		}
	}
	return primes
}
