// Package listops provides a solution to the listops exercise of the Go track in https://exercism.io
package listops

// IntList a lis of ints
type IntList []int

type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

// Foldl fold left
func (ints IntList) Foldl(f binFunc, initial int) int {
	if ints.Length() == 0 {
		return initial
	}
	i := f(initial, ints[0])
	return ints[1:].Foldl(f, i)
}

// Foldr fold right
func (ints IntList) Foldr(f binFunc, initial int) int {
	bf := func(x, y int) int { return f(y, x) }
	return ints.Reverse().Foldl(bf, initial)
}

// Filter keep elements for which the function returns true
func (ints IntList) Filter(f predFunc) IntList {
	switch {
	case ints.Length() == 0:
		return ints
	case f(ints[0]):
		return IntList{ints[0]}.Append(ints[1:].Filter(f))
	default:
		return ints[1:].Filter(f)
	}
}

// Length the list length
func (ints IntList) Length() int {
	length := 0
	for range ints {
		length++
	}
	return length
}

// Map every element using a function
func (ints IntList) Map(f unaryFunc) IntList {
	if ints.Length() == 0 {
		return ints
	}
	return IntList{f(ints[0])}.Append(ints[1:].Map(f))
}

// Reverse the list
func (ints IntList) Reverse() IntList {
	if ints.Length() == 0 {
		return ints
	}
	return ints[1:].Reverse().Append([]int{ints[0]})
}

// Append a list
func (ints IntList) Append(list IntList) IntList {
	totalLength := ints.Length() + list.Length()
	newInts := make(IntList, totalLength)
	for index := 0; index < ints.Length(); index++ {
		newInts[index] = ints[index]
	}
	for index := 0; index < list.Length(); index++ {
		newInts[ints.Length()+index] = list[index]
	}
	return newInts
}

// Concat the elements in all the lists
func (ints IntList) Concat(list []IntList) IntList {
	newInts := ints
	for _, li := range list {
		newInts = newInts.Append(li)
	}
	return newInts
}
