package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(&GameOfLife{}); err != nil {
		log.Fatal(err)
	}
}
