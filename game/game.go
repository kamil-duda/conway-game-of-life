package game

import (
	"bytes"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/conway"
	"golang.org/x/image/font/gofont/gomono"
)

type GameOfLife struct {
	universe *Universe
	canvas   *gameCanvas
}

type FpsCounter struct {
	frames uint
}

func (game *GameOfLife) Update() error {
	nextUniverse := game.universe.clone()
	for x := 0; x < config.LogicalWidth; x++ {
		for y := 0; y < config.LogicalHeight; y++ {
			neighbours := game.universe.liveNeighbours(x, y)
			if game.universe.isLive(x, y) {
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
	game.universe = nextUniverse
	return nil
}

func (game *GameOfLife) Draw(screen *ebiten.Image) {
	game.canvas.clear()
	for cell := range game.universe.cellsIter() {
		game.canvas.pixel(cell.x, cell.y)
	}
	game.canvas.drawOnto(screen)

	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		panic(err)
	}
	face := &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(0, 0)
	text.Draw(screen, "Conway's Game of Life", face, op)
}

func (game *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
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
	gameBuffer := newCanvas(sizeX, sizeY)
	return &GameOfLife{universe, gameBuffer}
}
