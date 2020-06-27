package utils

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	tStarted time.Time
	Elapsed  time.Duration
}

func (s *Stopwatch) Start() {
	s.tStarted = time.Now()
}

func (s *Stopwatch) Stop() {
	tEnded := time.Now()
	s.Elapsed = tEnded.Sub(s.tStarted)
}

var formats = [...]string{"ns", "µs", "ms", "s"}

func (s *Stopwatch) ElapsedReadable() string {
	tElapsed := s.Elapsed.Nanoseconds()
	format := formats[0]
	for i := 1; tElapsed >= 1000 && i < 3; i++ {
		format = formats[i]
		tElapsed /= 1000
	}

	return fmt.Sprintf("%d%s", tElapsed, format)
}
