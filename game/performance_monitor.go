package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type performanceMonitor struct {
	fpsCounter       *rateCounter
	simSpeedCounter  *rateCounter
	fpsRenderer      *draw.FpsRenderer
	simSpeedRenderer *draw.SimSpeedRenderer
}

func newPerformanceMonitor() *performanceMonitor {
	return &performanceMonitor{
		&rateCounter{},
		&rateCounter{},
		draw.NewFpsRenderer(),
		draw.NewSimSpeedRenderer(),
	}
}

func (p *performanceMonitor) fpsTick() {
	p.fpsCounter.tick()
}

func (p *performanceMonitor) simSpeedTick() {
	p.simSpeedCounter.tick()
}

func (p *performanceMonitor) draw(screen *ebiten.Image) {
	p.fpsRenderer.Draw(p.fpsCounter.rate, screen)
	p.simSpeedRenderer.Draw(p.fpsCounter.rate, screen)
}
