package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetWindowSize(256, 128)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&GameOfLife{}); err != nil {
		log.Fatal(err)
	}
}
