package game

import (
	"fmt"
	"time"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type performanceMonitor struct {
	fpsCounter      *rateCounter
	simSpeedCounter *rateCounter
	generationCount int
	populationCount int
	aliveCount      int
	deadCount       int
	activityLevel   string
	startTime       time.Time
	// Legacy renderers - will be removed
	fpsRenderer        *draw.MetricRenderer
	simSpeedRenderer   *draw.MetricRenderer
	generationRenderer *draw.MetricRenderer
	populationRenderer *draw.MetricRenderer
	// UI widgets
	fpsLabel      *widget.Label
	upsLabel      *widget.Label
	genLabel      *widget.Label
	popLabel      *widget.Label
	aliveLabel    *widget.Label
	deadLabel     *widget.Label
	activityLabel *widget.Label
	runtimeLabel  *widget.Label
}

func newPerformanceMonitor() *performanceMonitor {
	return &performanceMonitor{
		fpsCounter:         &rateCounter{},
		simSpeedCounter:    &rateCounter{},
		generationCount:    0,
		populationCount:    0,
		aliveCount:         0,
		deadCount:          0,
		activityLevel:      "Low",
		startTime:          time.Now(),
		fpsRenderer:        draw.NewFpsRenderer(),
		simSpeedRenderer:   draw.NewSimSpeedRenderer(),
		generationRenderer: draw.NewGenerationRenderer(),
		populationRenderer: draw.NewPopulationRenderer(),
		fpsLabel:           nil,
		upsLabel:           nil,
		genLabel:           nil,
		popLabel:           nil,
		aliveLabel:         nil,
		deadLabel:          nil,
		activityLabel:      nil,
		runtimeLabel:       nil,
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

func (p *performanceMonitor) setCellCounts(aliveCount, deadCount int) {
	p.aliveCount = aliveCount
	p.deadCount = deadCount

	// Calculate activity level based on population changes
	if aliveCount > p.populationCount*2 {
		p.activityLevel = "High"
	} else if aliveCount > p.populationCount {
		p.activityLevel = "Medium"
	} else {
		p.activityLevel = "Low"
	}
}

func (p *performanceMonitor) setMetricsWidgets(fpsLabel, upsLabel, genLabel, popLabel, aliveLabel, deadLabel, activityLabel, runtimeLabel *widget.Label) {
	p.fpsLabel = fpsLabel
	p.upsLabel = upsLabel
	p.genLabel = genLabel
	p.popLabel = popLabel
	p.aliveLabel = aliveLabel
	p.deadLabel = deadLabel
	p.activityLabel = activityLabel
	p.runtimeLabel = runtimeLabel
}

func (p *performanceMonitor) updateWidgetTexts() {
	// Row 1: FPS, UPS, Generation, Population (with live counters)
	if p.fpsLabel != nil {
		p.fpsLabel.Label = fmt.Sprintf("🎯 FPS: %d", p.fpsCounter.rate)
	}
	if p.upsLabel != nil {
		p.upsLabel.Label = fmt.Sprintf("⚡ UPS: %d", p.simSpeedCounter.rate)
	}
	if p.genLabel != nil {
		p.genLabel.Label = fmt.Sprintf("🧬 Generation: %d", p.generationCount)
	}
	if p.popLabel != nil {
		p.popLabel.Label = fmt.Sprintf("👥 Population: %d", p.populationCount)
	}

	// Row 2: Alive cells, Dead cells, Activity level, Runtime
	if p.aliveLabel != nil {
		p.aliveLabel.Label = fmt.Sprintf("📊 Alive: %d", p.aliveCount)
	}
	if p.deadLabel != nil {
		p.deadLabel.Label = fmt.Sprintf("💀 Dead: %d", p.deadCount)
	}
	if p.activityLabel != nil {
		p.activityLabel.Label = fmt.Sprintf("🔄 Activity: %s", p.activityLevel)
	}
	if p.runtimeLabel != nil {
		runtime := time.Since(p.startTime)
		minutes := int(runtime.Minutes())
		seconds := int(runtime.Seconds()) % 60
		p.runtimeLabel.Label = fmt.Sprintf("⏱️ Runtime: %d:%02d", minutes, seconds)
	}
}

func (p *performanceMonitor) resetGeneration() {
	p.generationCount = 0
}
