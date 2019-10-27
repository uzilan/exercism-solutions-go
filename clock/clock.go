// Package clock provides a solution to the Clock exercise of the Go track in https://exercism.io
package clock

import "fmt"

// Clock implements a clock that handles times without dates
type Clock struct {
	hours   int
	minutes int
}

// New creates a new clock with given hours and minutes.
// If the minutes or the hours are not inside the normal intervals, they are converted to do so
func New(hours int, minutes int) Clock {
	if minutes >= 60 {
		hours += minutes / 60
		minutes %= 60
	}

	for minutes < 0 {
		minutes += 60
		hours--
	}

	if hours >= 24 {
		hours %= 24
	}

	for hours < 0 {
		hours += 24
	}

	return Clock{
		hours,
		minutes,
	}
}

// String returns a textual representation of the clock in format hh:mm
func (clock Clock) String() string {
	var hours string
	var minutes string

	switch h := clock.hours; {
	case h == 0:
		hours = "00"
	case h < 10:
		hours = fmt.Sprintf("0%v", clock.hours)
	default:
		hours = fmt.Sprintf("%v", clock.hours)
	}

	switch m := clock.minutes; {
	case m == 0:
		minutes = "00"
	case m < 10:
		minutes = fmt.Sprintf("0%v", clock.minutes)
	default:
		minutes = fmt.Sprintf("%v", clock.minutes)
	}

	return fmt.Sprintf("%v:%v", hours, minutes)
}

// Add adds given minutes to the clock
func (clock Clock) Add(minutes int) Clock {
	return New(clock.hours, clock.minutes+minutes)
}

// Subtract subtracts given minutes from the clock
func (clock Clock) Subtract(minutes int) Clock {
	return New(clock.hours, clock.minutes-minutes)
}
