package game

import "time"

type fpsCounter struct {
	frames          uint
	lastCalculation time.Time
	fps             uint
}

func (f *fpsCounter) tick() {
	resetCounters := func() {
		f.frames = 0
		f.lastCalculation = time.Now()
	}

	// if more than 1 second passed - calculate fps with current values, do not add a new frame now
	switch interval := time.Since(f.lastCalculation); {
	case interval >= time.Second*2:
		f.fps = 0
		resetCounters()
	case interval >= time.Second:
		f.fps = f.frames
		resetCounters()
	}
	// add a new frame only after calculations are done
	f.frames++
}
