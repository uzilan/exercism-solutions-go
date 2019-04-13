// Package etl provides a solution to the Nth Prime exercise
// of the Go track in https://exercism.io
package prime

// Nth determines what the nth prime is, given a number.
func Nth(nth int) (int, bool) {
	if nth <= 0 {
		return -1, false
	}

	var primes = make([]int, 0)
	result := findNthPrime(&primes, 2, nth)

	return result, true
}

func findNthPrime(primes *[]int, curr int, nth int) int {
	primesFound := len(*primes)
	if primesFound == nth {
		return (*primes)[primesFound-1]
	}

	for _, p := range *primes {
		if curr%p == 0 {
			return findNthPrime(primes, curr+1, nth)
		}
	}

	*primes = append(*primes, curr)
	return findNthPrime(primes, curr+1, nth)
}
