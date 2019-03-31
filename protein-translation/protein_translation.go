// Package protein provides a solution to the Protein Translation exercise of the Go track in https://exercism.io
package protein

import "errors"

// ErrStop found a terminating codon
var ErrStop = errors.New("ErrStop")

// ErrInvalidBase found a wrong codon
var ErrInvalidBase = errors.New("ErrInvalidBase")

// FromCodon translates an RNA sequence into  a protein.
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	}
	return "", ErrInvalidBase
}

// FromRNA translates RNA sequences into proteins.
func FromRNA(input string) ([]string, error) {
	var proteins []string
	for i := 0; i <= len(input)-3; i += 3 {
		protein, err := FromCodon(input[i : i+3])
		if ErrStop == err {
			return proteins, nil
		} else if ErrInvalidBase == err {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
