// Package proverb provides a solution to the proverb exercise of the Go track in https://exercism.io
package proverb

import "fmt"

// Proverb generate a famous proverb given a list of inputs
func Proverb(rhyme []string) []string {
	var rows []string
	for index, word := range rhyme {
		if index == len(rhyme)-1 {
			lastRow := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			rows = append(rows, lastRow)
		} else {
			row := fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[index+1])
			rows = append(rows, row)
		}
	}
	return rows
}
