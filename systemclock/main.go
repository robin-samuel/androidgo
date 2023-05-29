package systemclock

import (
	"math/rand"
	"time"
)

type SystemClock struct {
	StartTime time.Time
}

func NewSystemClock() *SystemClock {
	START_TIME := time.Now()
	// Substract random time to avoid detection (2 hours up to 2 weeks)
	sub := time.Duration(rand.Int63n(60*60*24*14-60*60*24*2)+60*60*24*2) * time.Second
	START_TIME = START_TIME.Add(-sub)
	return &SystemClock{StartTime: START_TIME}
}

func (s *SystemClock) UptimeMillis() int64 {
	return time.Since(s.StartTime).Milliseconds()
}
