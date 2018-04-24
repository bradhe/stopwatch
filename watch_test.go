package stopwatch

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

const (
	expectedMilliseconds = (86400 * 1000)
)

func withNow(fn func() time.Time, callback func()) {
	oldNow := now
	defer func() {
		now = oldNow
	}()

	now = fn
	callback()
}

func withNowOffset(t time.Duration, callback func()) {
	fn := func() time.Time {
		return time.Now().Add(t)
	}

	withNow(fn, callback)
}

func TestStopWatchString(t *testing.T) {
	exp := `^30\.(\d+)ms$`
	rexp := regexp.MustCompile(exp)

	var watch Watch

	withNowOffset(-30*time.Millisecond, func() {
		watch = Start()
	})

	watch.Stop()

	// We're not millisecond accurate above, so...
	if !rexp.MatchString(watch.String()) {
		t.Fatalf("expected `%s` to match `%s`", watch, exp)
	}
}

func TestDeferring(t *testing.T) {
	exp := `^30m0\.\d+s$`
	rexp := regexp.MustCompile(exp)

	var called bool

	defer func() {
		if !called {
			t.Fatalf("failed to call defered function")
		}
	}()

	var watch Watch

	// Rewind the clock by 30 minutes so we have a realistic value to check this
	// against.
	withNowOffset(-30*time.Minute, func() {
		watch = Start()
	})

	defer watch.Timer(func(w Watch) {
		called = true

		if !rexp.MatchString(w.String()) {
			t.Fatalf("expected `%s` to match `%s`", watch, exp)
		}
	})
}

func ExampleWatch_Timer() {
	defer StartAt(time.Now().Add(-30 * time.Minute)).Timer(func(w Watch) {
		fmt.Printf("elapsed time: %d minutes", w.Minutes())
	})

	// Output:
	// elapsed time: 30 minutes
}
