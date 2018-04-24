package stopwatch

import (
	"fmt"
	"time"
)

// Function prototype for timers.
type TimerFunc func(Watch)

type Watch interface {
	fmt.Stringer

	// Timer calls a callback with the currently-measured time. This is useful for
	// deferring out of a function.
	Timer(fn TimerFunc)

	// Stops the watch based on the current wall-clock time.
	Stop() Watch

	// Starts the watch based on the current wall-clock time.
	Start() Watch

	// Milliseconds returns the elapsed duration in milliseconds.
	Milliseconds() time.Duration

	// Seconds returns the elapsed duration in seconds.
	Seconds() time.Duration

	// Minutes returns the elapsed duration in minutes.
	Minutes() time.Duration

	// Hours returns the elapsed duration in hours.
	Hours() time.Duration

	// Days returns the elapsed duration in days.
	Days() time.Duration
}

var now = func() time.Time {
	return time.Now()
}

type watch struct {
	start, stop time.Time
}

// Timer calls a callback with the currently-measured time. This is useful for
// deferring out of a function.
func (s *watch) Timer(fn TimerFunc) {
	fn(s.Stop())
}

// Stops the watch based on the current wall-clock time.
func (s *watch) Stop() Watch {
	s.stop = now()
	return s
}

// Starts the watch based on the current wall-clock time.
func (s *watch) Start() Watch {
	s.start = now()
	return s
}

// String returns a human-readable representation of the stopwatch's duration.
func (s *watch) String() string {
	// if the watch isn't stopped yet...
	if s.stop.IsZero() {
		return "0m0.00s"
	}

	return s.duration().String()
}

func (s *watch) duration() time.Duration {
	return s.stop.Sub(s.start)
}

// Milliseconds returns the elapsed duration in milliseconds.
func (s *watch) Milliseconds() time.Duration {
	return s.duration() / time.Millisecond
}

// Seconds returns the elapsed duration in seconds.
func (s *watch) Seconds() time.Duration {
	return s.duration() / time.Second
}

// Minutes returns the elapsed duration in minutes.
func (s *watch) Minutes() time.Duration {
	return s.duration() / time.Minute
}

// Hours returns the elapsed duration in hours.
func (s *watch) Hours() time.Duration {
	return s.duration() / time.Hour
}

// Days returns the elapsed duration in days.
func (s *watch) Days() time.Duration {
	return s.duration() / (24 * time.Hour)
}
