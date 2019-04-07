// Package anagram provides a solution to the Anagram exercise of the Go track in https://exercism.io
package anagram

import (
	"strings"
)

// Detect selects the correct sublist given a word and a list of possible anagrams
func Detect(subject string, candidates []string) []string {
	var result []string
	lowerSubject := strings.ToLower(subject)
	subjectMap := mapWord(lowerSubject)
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if lowerCandidate == lowerSubject {
			continue
		}
		if mapEquals(subjectMap, mapWord(lowerCandidate)) {
			result = append(result, candidate)
		}
	}
	return result
}

func mapWord(word string) map[rune]int {
	letterMap := map[rune]int{}
	for _, r := range word {
		letterMap[r]++
	}
	return letterMap
}

func mapEquals(mapA map[rune]int, mapB map[rune]int) bool {
	if len(mapA) != len(mapB) {
		return false
	}
	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}
	return true
}
