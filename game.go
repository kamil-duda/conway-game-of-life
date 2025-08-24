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
	draw.Pixel(0, 0, screen)
	draw.Pixel(5, 5, screen)
	draw.Pixel(10, 10, screen)
}

func (g *GameOfLife) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 32, 24
}
