// Package triangle provides a solution to the triangle exercise of the Go track in https://exercism.io
package triangle

import "math"

// Kind defines the type of the triangle
type Kind int

// different types of triangle
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral - all sides are equal
	Iso             // isosceles - two sides are equal
	Sca             // scalene - al sides are unique
)

// KindFromSides determines if a triangle legal, or is equilateral, isosceles, or scalene
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isLegalTriangle(a, b, c):
		return NaT
	case a == b && b == c:
		return Equ
	case a != b && a != c && b != c:
		return Sca
	default:
		return Iso
	}
}

// isLegalTriangle determines if a triangle is legal, that is that every side is legal,
// and that the sum of the sides is not smaller than twice the largest side
func isLegalTriangle(a, b, c float64) bool {
	var sides = []float64{a, b, c}

	for _, side := range sides {
		if !isLegalSide(side) {
			return false
		}
	}

	var max = math.Max(a, math.Max(b, c))
	if max*2 > a+b+c {
		return false
	}

	return true
}

// isLegalSide determines if a triangle side is legal, that is that it is a number,
// is not infinity and is bigger than zero
func isLegalSide(side float64) bool {
	return !math.IsNaN(side) && !math.IsInf(side, 0) && side > 0
}
