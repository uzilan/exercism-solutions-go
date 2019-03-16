// An solution to the two-fer exercise of the Go track in https://exercism.io
package twofer

import "fmt"

// return a line about sharing something with the given name.
// If the name is empty, use "you" instead
func ShareWith(name string) string {
	var who = name
	if len(name) == 0 {
		who = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", who)
}
