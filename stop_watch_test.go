package stopwatch

import (
	"time"
	"testing"
)

const (
	expectedMilliseconds = (86400 * 1000)
)

func TestStopWatchCalculatesMilliseconds(t *testing.T) {
	start := time.Unix(int64(time.Now().Unix() - 86400), 0)
	watch := Stop(start)

	// We're not millisecond accurate above, so...
	if (watch.Milliseconds() - expectedMilliseconds) > 1000 {
		t.Logf("Expected milliseconds to be %d, got %d", expectedMilliseconds, watch.Milliseconds())
		t.Fail()
	}
}
