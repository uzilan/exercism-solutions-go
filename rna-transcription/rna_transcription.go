// Package strand provides a solution to the rna transcription exercise of the Go track in https://exercism.io
package strand

// ToRNA transcripts an DNA strand into a RNA strand
func ToRNA(dna string) string {
	var rna = ""
	for _, char := range dna {
		switch char {
		case 'C':
			rna += "G"
		case 'G':
			rna += "C"
		case 'T':
			rna += "A"
		case 'A':
			rna += "U"
		}
	}
	return rna
}
