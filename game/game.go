package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/conway"
)

type GameOfLife struct {
	universe *universe
	canvas   *gameCanvas
}

type fpsCounter struct {
	frames uint
}

func (g *GameOfLife) Update() error {
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
	return nil
}

func (g *GameOfLife) Draw(screen *ebiten.Image) {
	g.canvas.clear()
	for cell := range g.universe.cellsIter() {
		g.canvas.draw(cell)
	}
	g.canvas.drawOnto(screen)
}

func (g *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func NewRandom(sizeX, sizeY int) *GameOfLife {
	universe := newUniverse()
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			if rand.Intn(2) == 1 {
				universe.setLive(x, y)
			}
		}
	}
	gameBuffer := newCanvas(sizeX, sizeY)
	return &GameOfLife{universe, gameBuffer}
}
