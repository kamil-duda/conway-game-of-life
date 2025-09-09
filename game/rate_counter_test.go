package game

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestFreshRateCounter(t *testing.T) {
	t.Parallel()
	counter := rateCounter{}

	gotTicks := counter.ticks
	wantTicks := uint(0)
	if gotTicks != wantTicks {
		t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}
	gotRate := counter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
	}
}

func TestFreshSingleTick(t *testing.T) {
	t.Parallel()
	counter := rateCounter{}
	counter.tick()

	gotTicks := counter.ticks
	wantTicks := uint(1)
	if gotTicks != wantTicks {
		t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}
	gotRate := counter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
	}
}

func TestFreshTwoTicks(t *testing.T) {
	t.Parallel()
	counter := rateCounter{}
	counter.tick()
	counter.tick()

	gotTicks := counter.ticks
	wantTicks := uint(2)
	if gotTicks != wantTicks {
		t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}
	gotRate := counter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
	}
}

func Test1Second60Rate(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := rateCounter{}
		for i := 0; i < 60; i++ {
			counter.tick()
		}

		time.Sleep(time.Second * 1)
		synctest.Wait()
		// force calculation after 1 second is passed
		counter.tick()

		gotTicks := counter.ticks
		wantTicks := uint(1)
		gotRate := counter.rate
		wantRate := uint(60)

		// 1 new frame
		if gotTicks != wantTicks {
			t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
		}

		// 60 ticks from the last period
		if gotRate != wantRate {
			t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
		}
	})
}

func Test2Seconds60Rate(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := rateCounter{}
		for i := 0; i < 60; i++ {
			counter.tick()
		}
		time.Sleep(time.Second * 1)
		synctest.Wait()
		for i := 0; i < 60; i++ {
			counter.tick()
		}
		time.Sleep(time.Second * 1)
		synctest.Wait()
		// force calculation after more than 1 second is passed
		counter.tick()

		gotTicks := counter.ticks
		wantTicks := uint(1)
		gotRate := counter.rate
		wantRate := uint(60)

		// 1 new frame
		if gotTicks != wantTicks {
			t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
		}

		if gotRate != wantRate {
			t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
		}
	})
}

func TestLessThan1Rate(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := rateCounter{}
		counter.tick()

		time.Sleep(time.Second * 2)
		synctest.Wait()
		// force calculation after 2 seconds is passed
		counter.tick()

		gotTicks := counter.ticks
		wantTicks := uint(1)
		gotRate := counter.rate
		wantRate := uint(0)

		// 1 new frame
		if gotTicks != wantTicks {
			t.Errorf("rateCounter.ticks = %d, want %d", gotTicks, wantTicks)
		}

		if gotRate != wantRate {
			t.Errorf("rateCounter.rate = %d, want %d", gotRate, wantRate)
		}
	})
}
