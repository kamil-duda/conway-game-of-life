package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/conway"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type GameOfLife struct {
	universe *Universe
}

func (g *GameOfLife) Update() error {
	nextUniverse := g.universe.clone()
	for x := 0; x < config.LogicalWidth; x++ {
		for y := 0; y < config.Height; y++ {
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
	for cell := range g.universe.cellsIter() {
		draw.Pixel(cell.x, cell.y, screen)
	}
	draw.DebugBackground(screen)
}

func (g *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.LogicalWidth, config.LogicalHeight
}

func NewRandom(sizeX, sizeY int) *GameOfLife {
	universe := New()
	for x := 0; x < sizeX; x++ {
		for y := 0; y < sizeY; y++ {
			if rand.Intn(2) == 1 {
				universe.setLive(x, y)
			}
		}
	}
	return &GameOfLife{universe}
}
