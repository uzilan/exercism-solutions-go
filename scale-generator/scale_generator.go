// Package scale provides a solution to the scale generator exercise of the Go track in https://exercism.io
package scale

import "strings"

// Step is an interval between two tones
type Step int

// different steps are used in different scales
const (
	Half      Step = 1
	Full           = 2
	Augmented      = 3
)

var sharpsScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatsScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var doubleSharpsScales = append(sharpsScale, sharpsScale...)
var doubleFlatsScales = append(flatsScale, flatsScale...)

// Scale generates the musical scale starting with the tonic and following the specified interval pattern
// given a tonic, or starting note, and a set of intervals,
func Scale(tonic string, interval string) []string {
	return helper(isSharps(tonic), strings.Title(tonic), interval, []string{strings.Title(tonic)})
}

func helper(sharps bool, note string, interval string, result []string) []string {
	var step Step
	switch {
	case len(interval) == 0, interval[0] == 'm':
		step = Half
	case interval[0] == 'M':
		step = Full
	case interval[0] == 'A':
		step = Augmented
	}
	nextNote := getNextNote(sharps, note, step)

	if result[0] == nextNote {
		return result
	}

	return helper(sharps, nextNote, getNextInterval(interval), append(result, nextNote))
}

func getNextInterval(interval string) string {
	if len(interval) == 0 {
		return ""
	}
	return interval[1:]
}

func isSharps(tonic string) bool {
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		return true
	}
	return false
}

func getNextNote(sharps bool, tonic string, step Step) string {
	if sharps {
		return searchNext(sharpsScale, doubleSharpsScales, tonic, step)
	}
	return searchNext(flatsScale, doubleFlatsScales, tonic, step)
}

func searchNext(scale []string, doubleScale []string, tonic string, step Step) string {
	for index, current := range scale {
		if current == tonic {
			return doubleScale[index+int(step)]
		}
	}
	return ""
}
