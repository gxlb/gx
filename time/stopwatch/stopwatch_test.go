package stopwatch

import (
	"testing"
	"time"
)

func TestStopwatch(t *testing.T) {
	var d WatchDuration
	d.Init("t", time.Second, 2)
	println(d.String())
}
