package game

import (
	"math/rand"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/conway"
	"github.com/kamil-duda/conway-game-of-life/ui"
)

type GameOfLife struct {
	ui                 *ebitenui.UI
	universe           *universe
	canvas             *gameCanvas
	performanceMonitor *performanceMonitor
	// Game state
	isPaused      bool
	stepRequested bool
	// Metrics widgets
	fpsLabel      *widget.Label
	upsLabel      *widget.Label
	genLabel      *widget.Label
	popLabel      *widget.Label
	aliveLabel    *widget.Label
	deadLabel     *widget.Label
	activityLabel *widget.Label
	runtimeLabel  *widget.Label
}

func (g *GameOfLife) Update() error {
	g.ui.Update()

	// Only advance simulation if not paused or if single step requested
	if !g.isPaused || g.stepRequested {
		g.stepRequested = false // Reset step flag

		nextUniverse := g.universe.clone()
		for x := 0; x < config.LogicalWidth; x++ {
			for y := 0; y < config.LogicalHeight; y++ {
				neighbours := g.universe.liveNeighbours(x, y)
				if g.universe.isLive(x, y) {
					if !conway.LiveCellSurvives(neighbours) {
						nextUniverse.setDead(x, y)
					}
				} else {
					if conway.DeadCellRevives(neighbours) {
						nextUniverse.setLive(x, y)
					}
				}
			}
		}
		g.universe = nextUniverse
		g.performanceMonitor.tickGeneration()
		g.performanceMonitor.simSpeedTick()
	}

	// Always update metrics regardless of pause state
	populationCount := g.universe.population()
	aliveCount := populationCount
	deadCount := (config.LogicalWidth * config.LogicalHeight) - populationCount

	g.performanceMonitor.setPopulationCount(populationCount)
	g.performanceMonitor.setCellCounts(aliveCount, deadCount)

	return nil
}

func (g *GameOfLife) Draw(screen *ebiten.Image) {
	g.canvas.clear()
	for cell := range g.universe.cellsIter() {
		g.canvas.draw(cell)
	}
	g.canvas.drawOnto(screen)
	g.performanceMonitor.updateWidgetTexts()
	g.ui.Draw(screen)
	g.performanceMonitor.fpsTick()
}

func (g *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (g *GameOfLife) resetSimulation() {
	g.generateNewRandomUniverse()
	g.performanceMonitor.resetGeneration()
}

func (g *GameOfLife) togglePause() {
	g.isPaused = !g.isPaused
}

func (g *GameOfLife) requestStep() {
	if g.isPaused {
		g.stepRequested = true
	}
}

func (g *GameOfLife) isPausedState() bool {
	return g.isPaused
}

func (g *GameOfLife) generateNewRandomUniverse() {
	newUniverse := newUniverse()
	for x := 0; x < config.LogicalWidth; x++ {
		for y := 0; y < config.LogicalHeight; y++ {
			if rand.Intn(2) == 1 {
				newUniverse.setLive(x, y)
			}
		}
	}
	g.universe = newUniverse
	g.performanceMonitor.setPopulationCount(newUniverse.population())
}

func NewRandomGame(sizeX, sizeY int) *GameOfLife {
	universe := newUniverse()
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			if rand.Intn(2) == 1 {
				universe.setLive(x, y)
			}
		}
	}
	gameBuffer := newCanvas(sizeX, sizeY)

	game := &GameOfLife{
		universe:           universe,
		canvas:             gameBuffer,
		performanceMonitor: newPerformanceMonitor(),
	}

	rootUiContainer, metricsWidgets := ui.InitializeUI(
		game.resetSimulation,
		game.togglePause,
		game.requestStep,
		func() { /* TODO: implement pattern loading */ },
		func() { /* TODO: implement settings dialog */ },
	)
	game.ui = &ebitenui.UI{Container: rootUiContainer}
	game.fpsLabel = metricsWidgets.FpsLabel
	game.upsLabel = metricsWidgets.UpsLabel
	game.genLabel = metricsWidgets.GenLabel
	game.popLabel = metricsWidgets.PopLabel
	game.aliveLabel = metricsWidgets.AliveLabel
	game.deadLabel = metricsWidgets.DeadLabel
	game.activityLabel = metricsWidgets.ActivityLabel
	game.runtimeLabel = metricsWidgets.RuntimeLabel
	game.performanceMonitor.setMetricsWidgets(
		metricsWidgets.FpsLabel, metricsWidgets.UpsLabel, metricsWidgets.GenLabel, metricsWidgets.PopLabel,
		metricsWidgets.AliveLabel, metricsWidgets.DeadLabel, metricsWidgets.ActivityLabel, metricsWidgets.RuntimeLabel)

	return game
}
