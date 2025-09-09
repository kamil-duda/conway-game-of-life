// Package game implements game mechanics and utilities
package game

import "time"

// rateCounter tracks and calculates rate of events (like frames per second)
// by counting ticks over time periods.
type rateCounter struct {
	ticks           uint
	lastCalculation time.Time
	rate            uint
}

// tick registers a new event and updates the rate if enough time has passed.
//
// The rate is calculated when at least 1 second passes since the last calculation.
// If more than 2 seconds pass, the rate is reset to 0.
func (f *rateCounter) tick() {
	resetCounters := func() {
		f.ticks = 0
		f.lastCalculation = time.Now()
	}

	// if more than 1 second passed - calculate fps with current values, do not add a new frame now
	switch interval := time.Since(f.lastCalculation); {
	case interval >= time.Second*2:
		f.rate = 0
		resetCounters()
	case interval >= time.Second:
		f.rate = f.ticks
		resetCounters()
	}
	// add a new frame only after calculations are done
	f.ticks++
}
