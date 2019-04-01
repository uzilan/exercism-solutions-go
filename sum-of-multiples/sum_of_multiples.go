// Package summultiples provides a solution to the Sum Of Multiples exercise of the Go track in https://exercism.io
package summultiples

//  SumMultiples finds the sum of all the unique multiples of particular numbers up to but not including a given number
func SumMultiples(limit int, devisors ...int) int {
	multipliers := map[int]bool{}
	for _, devisor := range devisors {
		if devisor != 0 {
			for i := 1; i < limit; i++ {
				if i%devisor == 0 && !multipliers[i] {
					multipliers[i] = true
				}
			}
		}
	}

	sum := 0
	for number := range multipliers {
		sum += number
	}
	return sum
}
