package stopwatch

import (
	"regexp"
	"testing"
	"time"
)

const (
	expectedMilliseconds = (86400 * 1000)
)

func TestStopWatchCalculatesMilliseconds(t *testing.T) {
	start := time.Unix(int64(time.Now().Unix()-86400), 0)
	watch := Stop(start)

	// We're not millisecond accurate above, so...
	if (watch.Milliseconds() - expectedMilliseconds) > 1000 {
		t.Logf("Expected milliseconds to be %d, got %d", expectedMilliseconds, watch.Milliseconds())
		t.Fail()
	}
}

func TestStopWatchString(t *testing.T) {
	exp := `^30\.(\d+)ms$`
	rexp := regexp.MustCompile(exp)

	start := time.Now().Add(-30 * time.Millisecond)
	watch := Stop(start)

	// We're not millisecond accurate above, so...
	if !rexp.MatchString(watch.String()) {
		t.Logf("Expected string match `%s` got %s", exp, watch.String())
		t.Fail()
	}
}
