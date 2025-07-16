package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type GameOfLife struct {
}

func (g *GameOfLife) Update() error {
	return nil
}

func (g *GameOfLife) Draw(screen *ebiten.Image) {
	draw.DrawSquare(screen)
}

func (g *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
