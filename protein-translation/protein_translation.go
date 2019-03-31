// Package protein provides a solution to the Protein Translation exercise of the Go track in https://exercism.io
package protein

// ErrStop ...
var ErrStop error

// ErrInvalidBase ...
var ErrInvalidBase error

// FromCodon translates RNA sequences into proteins.
func FromCodon(input string) (string, error) {
	protein := ""
	for i := 0; i < len(input)/3; i += 3 {
		protein += rnaToProtein(input[i*3 : i*3+3])
	}

	return protein, nil
}

// FromRNA ...
func FromRNA(input string) (string, error) {
	return "", nil
}

func rnaToProtein(rna string) string {
	switch rna {
	case "AUG":
		return "Methionine"
	case "UUU", "UUC":
		return "Phenylalanine"
	case "UUA", "UUG":
		return "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine"
	case "UAU", "UAC":
		return "Tyrosine"
	case "UGU", "UGC":
		return "Cysteine"
	case "UGG":
		return "Tryptophan"
	case "UAA", "UAG", "UGA":
		return "STOP"
	}
	return ""
}
