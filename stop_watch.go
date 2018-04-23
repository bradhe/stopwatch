package stopwatch

import (
	"fmt"
	"time"
)

type StopWatch struct {
	start, stop time.Time
}

func (s *StopWatch) Milliseconds() uint32 {
	return uint32(s.stop.Sub(s.start) / time.Millisecond)
}

func (s *StopWatch) String() string {
	return fmt.Sprintf("%s", s.stop.Sub(s.start))
}
