package tasks

import (
	"time"
)

const defaultTick = time.Second

type TimeWheel struct {
	tick time.Duration
}

func NewDefaultTimeWheel() *TimeWheel {
	return NewTimeWheel(0)
}

func NewTimeWheel(tick time.Duration) *TimeWheel {
	if tick <= 0 {
		tick = defaultTick
	}
	return &TimeWheel{tick: tick}
}
