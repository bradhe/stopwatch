package stopwatch

import (
	"time"
)

// Start starts a new Watch for you.
func Start() Watch {
	watch := &watch{time.Time{}, time.Time{}}
	return watch.Start()
}

// StartAt starts a new Watch for you at the time supplied. Mostly useful for
// testing, but could be used to capture other logic within an app or service.
func StartAt(t time.Time) Watch {
	return &watch{t, time.Time{}}
}
