// Package strain provides a solution to the strain exercise of the Go track in https://exercism.io
package strain

// Ints is a list of integers
type Ints []int

// Strings is a list of strings
type Strings []string

// Lists is a list of Ints
type Lists []Ints

// Keep returns a new collection containing those elements where the predicate is true
func (ints Ints) Keep(predicate func(x int) bool) Ints {
	if ints == nil {
		return nil
	}

	newInts := Ints{}
	for _, item := range ints {
		if predicate(item) {
			newInts = append(newInts, item)
		}
	}
	return newInts
}

// Discard returns a new collection containing those elements where the predicate is false
func (ints Ints) Discard(predicate func(x int) bool) Ints {
	return ints.Keep(func(x int) bool { return !predicate(x) })
}

// Keep returns a new collection containing those elements where the predicate is true
func (strings Strings) Keep(predicate func(s string) bool) Strings {
	if strings == nil {
		return nil
	}

	newStrings := Strings{}
	for _, item := range strings {
		if predicate(item) {
			newStrings = append(newStrings, item)
		}
	}
	return newStrings
}

// Discard returns a new collection containing those elements where the predicate is false
func (strings Strings) Discard(predicate func(s string) bool) Strings {
	return strings.Keep(func(s string) bool { return !predicate(s) })
}

// Keep returns a new collection containing those elements where the predicate is true
func (lists Lists) Keep(predicate func(s []int) bool) Lists {
	if lists == nil {
		return nil
	}

	newLists := Lists{}
	for _, item := range lists {
		if predicate(item) {
			newLists = append(newLists, item)
		}
	}
	return newLists
}

// Discard returns a new collection containing those elements where the predicate is false
func (lists Lists) Discard(predicate func(s []int) bool) Lists {
	return lists.Keep(func(s []int) bool { return !predicate(s) })
}
