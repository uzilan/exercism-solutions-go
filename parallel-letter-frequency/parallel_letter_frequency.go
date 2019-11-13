// Package letter provides a solution to the Parallel Letter Frequency exercise of the Go track in https://exercism.io
package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of letters in texts using parallel computation
func ConcurrentFrequency(rows []string) FreqMap {
	freqMap := FreqMap{}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(rows))
	for _, row := range rows {
		go func(s string) {
			defer waitGroup.Done()
			for k, v := range Frequency(s) {
				freqMap[k] += v
			}
		}(row)
	}
	waitGroup.Wait()
	return freqMap
}
