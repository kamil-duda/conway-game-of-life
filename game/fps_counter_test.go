package game

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestFreshFpsCounter(t *testing.T) {
	t.Parallel()
	counter := fpsCounter{}

	gotFrames := counter.frames
	wantFrames := uint(0)
	if gotFrames != wantFrames {
		t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
	}
	gotFps := counter.fps
	wantFps := uint(0)
	if gotFps != wantFps {
		t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
	}
}

func TestFreshSingleTick(t *testing.T) {
	t.Parallel()
	counter := fpsCounter{}
	counter.tick()

	gotFrames := counter.frames
	wantFrames := uint(1)
	if gotFrames != wantFrames {
		t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
	}
	gotFps := counter.fps
	wantFps := uint(0)
	if gotFps != wantFps {
		t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
	}
}

func TestFreshTwoTicks(t *testing.T) {
	t.Parallel()
	counter := fpsCounter{}
	counter.tick()
	counter.tick()

	gotFrames := counter.frames
	wantFrames := uint(2)
	if gotFrames != wantFrames {
		t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
	}
	gotFps := counter.fps
	wantFps := uint(0)
	if gotFps != wantFps {
		t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
	}
}

func Test1Second60Fps(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := fpsCounter{}
		for i := 0; i < 60; i++ {
			counter.tick()
		}

		time.Sleep(time.Second * 1)
		synctest.Wait()
		// force calculation after 1 second is passed
		counter.tick()

		gotFrames := counter.frames
		wantFrames := uint(1)
		gotFps := counter.fps
		wantFps := uint(60)

		// 1 new frame
		if gotFrames != wantFrames {
			t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
		}

		// 60 frames from the last period
		if gotFps != wantFps {
			t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
		}
	})
}

func Test2Seconds60Fps(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := fpsCounter{}
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

		gotFrames := counter.frames
		wantFrames := uint(1)
		gotFps := counter.fps
		wantFps := uint(60)

		// 1 new frame
		if gotFrames != wantFrames {
			t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
		}

		if gotFps != wantFps {
			t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
		}
	})
}

func TestLessThan1Fps(t *testing.T) {
	t.Parallel()
	synctest.Test(t, func(t *testing.T) {
		counter := fpsCounter{}
		counter.tick()

		time.Sleep(time.Second * 2)
		synctest.Wait()
		// force calculation after 2 seconds is passed
		counter.tick()

		gotFrames := counter.frames
		wantFrames := uint(1)
		gotFps := counter.fps
		wantFps := uint(0)

		// 1 new frame
		if gotFrames != wantFrames {
			t.Errorf("fpsCounter.frames = %d, want %d", gotFrames, wantFrames)
		}

		if gotFps != wantFps {
			t.Errorf("fpsCounter.fps = %d, want %d", gotFps, wantFps)
		}
	})
}
