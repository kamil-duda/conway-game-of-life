package game

import "time"

type fpsCounter struct {
	frames          uint
	lastCalculation time.Time
	fps             uint
}

func (fps *fpsCounter) tick() {
	resetCounters := func() {
		fps.frames = 0
		fps.lastCalculation = time.Now()
	}

	// if more than 1 second passed - calculate fps with current values, do not add a new frame now
	switch intervalSeconds := time.Since(fps.lastCalculation).Seconds(); {
	case intervalSeconds >= 2:
		fps.fps = 0
		resetCounters()
	case intervalSeconds >= 1:
		fps.fps = fps.frames
		resetCounters()
	}
	// add a new frame only after calculations are done
	fps.frames++
}
