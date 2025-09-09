package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type performanceMonitor struct {
	fpsCounter         *rateCounter
	simSpeedCounter    *rateCounter
	generationCount    int
	populationCount    int
	fpsRenderer        *draw.MetricRenderer
	simSpeedRenderer   *draw.MetricRenderer
	generationRenderer *draw.MetricRenderer
	populationRenderer *draw.MetricRenderer
}

func newPerformanceMonitor() *performanceMonitor {
	return &performanceMonitor{
		&rateCounter{},
		&rateCounter{},
		0,
		0,
		draw.NewFpsRenderer(),
		draw.NewSimSpeedRenderer(),
		draw.NewGenerationRenderer(),
		draw.NewPopulationRenderer(),
	}
}

func (p *performanceMonitor) draw(screen *ebiten.Image) {
	p.fpsRenderer.Draw(p.fpsCounter.rate, screen)
	p.simSpeedRenderer.Draw(p.simSpeedCounter.rate, screen)
	p.generationRenderer.Draw(p.generationCount, screen)
	p.populationRenderer.Draw(p.populationCount, screen)
}

func (p *performanceMonitor) fpsTick() {
	p.fpsCounter.tick()
}

func (p *performanceMonitor) simSpeedTick() {
	p.simSpeedCounter.tick()
}

func (p *performanceMonitor) tickGeneration() {
	p.generationCount++
}

func (p *performanceMonitor) setPopulationCount(populationCount int) {
	p.populationCount = populationCount
}
