// Package clock provides a solution to the Clock exercise of the Go track in https://exercism.io
package clock

import "fmt"

// Clock implements a clock that handles times without dates
type Clock struct {
	minutes int
}

const minutesPerDay = 1440

// New creates a new clock with given hours and minutes.
func New(hours int, minutes int) Clock {
	minutes += 60 * hours
	minutes %= minutesPerDay
	if minutes < 0 {
		minutes += minutesPerDay
	}
	return Clock{minutes}
}

// String returns a textual representation of the clock in format hh:mm
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.minutes/60, clock.minutes%60)
}

// Add adds given minutes to the clock
func (clock Clock) Add(minutes int) Clock {
	return New(0, clock.minutes+minutes)
}

// Subtract subtracts given minutes from the clock
func (clock Clock) Subtract(minutes int) Clock {
	return New(0, clock.minutes-minutes)
}
