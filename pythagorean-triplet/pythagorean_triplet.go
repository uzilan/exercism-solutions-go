// Package pythagorean provides a solution to the Pythagorean Triplet exercise of the Go track in https://exercism.io
package pythagorean

import "sort"

// Triplet is a pythagorean triplet
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	return findTrippleats(min, max, []func(x, y, z int) bool{
		func(x, y, z int) bool { return isTriplet(x, y, z) },
	})
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	return findTrippleats(1, p, []func(x, y, z int) bool{
		func(x, y, z int) bool { return isTriplet(x, y, z) },
		func(x, y, z int) bool { return isEqualsPerimeter(x, y, z, p) },
	})
}

func findTrippleats(min, max int, fs []func(x, y, z int) bool) []Triplet {
	var triplets []Triplet
	for z := min; z <= max; z++ {
		for y := min; y <= z; y++ {
			for x := min; x <= y; x++ {
				ok := true
				for _, f := range fs {
					if !f(x, y, z) {
						ok = false
					}
				}
				if ok {
					triplets = append(triplets, [3]int{x, y, z})
				}
			}
		}
	}
	sort.Slice(triplets, func(i, j int) bool {
		return triplets[i][0] < triplets[j][0]
	})
	return triplets
}

func isTriplet(x, y, z int) bool {
	return square(x)+square(y) == square(z)
}

func isEqualsPerimeter(x, y, z, p int) bool {
	return x+y+z == p
}

func square(i int) int {
	return i * i
}
