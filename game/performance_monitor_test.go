package game

import "testing"

func TestNewPerformanceMonitor(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()

	// Test that fps counter has zero values
	gotFpsTicks := monitor.fpsCounter.ticks
	wantFpsTicks := uint(0)
	if gotFpsTicks != wantFpsTicks {
		t.Errorf("newPerformanceMonitor().fpsCounter.ticks = %d, want %d", gotFpsTicks, wantFpsTicks)
	}

	gotFpsRate := monitor.fpsCounter.rate
	wantFpsRate := uint(0)
	if gotFpsRate != wantFpsRate {
		t.Errorf("newPerformanceMonitor().fpsCounter.rate = %d, want %d", gotFpsRate, wantFpsRate)
	}

	// Test that sim speed counter has zero values
	gotSimSpeedTicks := monitor.simSpeedCounter.ticks
	wantSimSpeedTicks := uint(0)
	if gotSimSpeedTicks != wantSimSpeedTicks {
		t.Errorf("newPerformanceMonitor().simSpeedCounter.ticks = %d, want %d", gotSimSpeedTicks, wantSimSpeedTicks)
	}

	gotSimSpeedRate := monitor.simSpeedCounter.rate
	wantSimSpeedRate := uint(0)
	if gotSimSpeedRate != wantSimSpeedRate {
		t.Errorf("newPerformanceMonitor().simSpeedCounter.rate = %d, want %d", gotSimSpeedRate, wantSimSpeedRate)
	}

	// Test that renderers are initialized
	if monitor.fpsRenderer == nil {
		t.Error("newPerformanceMonitor().fpsRenderer is nil, expected initialized renderer")
	}

	if monitor.simSpeedRenderer == nil {
		t.Error("newPerformanceMonitor().simSpeedRenderer is nil, expected initialized renderer")
	}
}

func TestFpsTickSingle(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()
	monitor.fpsTick()

	gotTicks := monitor.fpsCounter.ticks
	wantTicks := uint(1)
	if gotTicks != wantTicks {
		t.Errorf("after fpsTick(), fpsCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}

	gotRate := monitor.fpsCounter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("after fpsTick(), fpsCounter.rate = %d, want %d", gotRate, wantRate)
	}

	gotSimTicks := monitor.simSpeedCounter.ticks
	wantSimTicks := uint(0)
	if gotSimTicks != wantSimTicks {
		t.Errorf("after fpsTick(), simSpeedCounter.ticks = %d, want %d", gotSimTicks, wantSimTicks)
	}
}

func TestFpsTickMultiple(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()

	for i := 0; i < 5; i++ {
		monitor.fpsTick()
	}

	gotTicks := monitor.fpsCounter.ticks
	wantTicks := uint(5)
	if gotTicks != wantTicks {
		t.Errorf("after 5 fpsTick() calls, fpsCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}

	gotRate := monitor.fpsCounter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("after 5 fpsTick() calls, fpsCounter.rate = %d, want %d", gotRate, wantRate)
	}
}

func TestSimSpeedTickSingle(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()
	monitor.simSpeedTick()

	gotTicks := monitor.simSpeedCounter.ticks
	wantTicks := uint(1)
	if gotTicks != wantTicks {
		t.Errorf("after simSpeedTick(), simSpeedCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}

	gotRate := monitor.simSpeedCounter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("after simSpeedTick(), simSpeedCounter.rate = %d, want %d", gotRate, wantRate)
	}

	gotFpsTicks := monitor.fpsCounter.ticks
	wantFpsTicks := uint(0)
	if gotFpsTicks != wantFpsTicks {
		t.Errorf("after simSpeedTick(), fpsCounter.ticks = %d, want %d", gotFpsTicks, wantFpsTicks)
	}
}

func TestSimSpeedTickMultiple(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()
	for i := 0; i < 3; i++ {
		monitor.simSpeedTick()
	}

	gotTicks := monitor.simSpeedCounter.ticks
	wantTicks := uint(3)
	if gotTicks != wantTicks {
		t.Errorf("after 3 simSpeedTick() calls, simSpeedCounter.ticks = %d, want %d", gotTicks, wantTicks)
	}

	gotRate := monitor.simSpeedCounter.rate
	wantRate := uint(0)
	if gotRate != wantRate {
		t.Errorf("after 3 simSpeedTick() calls, simSpeedCounter.rate = %d, want %d", gotRate, wantRate)
	}
}

func TestBothCountersIndependent(t *testing.T) {
	t.Parallel()
	monitor := newPerformanceMonitor()
	monitor.fpsTick()
	monitor.fpsTick()
	monitor.simSpeedTick()

	gotFpsTicks := monitor.fpsCounter.ticks
	wantFpsTicks := uint(2)
	if gotFpsTicks != wantFpsTicks {
		t.Errorf("fpsCounter.ticks = %d, want %d", gotFpsTicks, wantFpsTicks)
	}

	gotSimTicks := monitor.simSpeedCounter.ticks
	wantSimTicks := uint(1)
	if gotSimTicks != wantSimTicks {
		t.Errorf("simSpeedCounter.ticks = %d, want %d", gotSimTicks, wantSimTicks)
	}
}
