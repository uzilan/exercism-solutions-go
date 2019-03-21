// Package gigasecond provides a solution to the gigasecond exercise of the Go track in https://exercism.io
package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
