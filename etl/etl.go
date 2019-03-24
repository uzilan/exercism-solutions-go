// Package etl provides a solution to the etl exercise of the Go track in https://exercism.io
package etl

import "strings"

// Transform transforms a legacy map structure into a new
func Transform(legacyMap map[int][]string) map[string]int {
	newMap := map[string]int{}

	for value, letters := range legacyMap {
		for _, letter := range letters {
			newMap[strings.ToLower(letter)] = value
		}
	}

	return newMap
}
