package stopwatch

import (
	"time"
)

type StopWatch struct {
	start, stop time.Time
}

func (self *StopWatch) Milliseconds() uint64 {
	return uint64(self.stop.Sub(self.start) / time.Millisecond)
}
