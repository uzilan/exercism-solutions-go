// Package dna provides a solution to the nucleotide exercise of the Go track in https://exercism.io
package dna

import (
	"errors"
	"regexp"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	invalidNucleotides, _ := regexp.MatchString("[^ACGT]", string(dna))
	if len(dna) > 0 && invalidNucleotides {
		return nil, errors.New("Found invalid nucleotides in the strand")
	}

	var histogram = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, c := range dna {
		histogram[c]++
	}

	return histogram, nil
}
