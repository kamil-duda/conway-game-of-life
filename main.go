package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/game"
)

func main() {
	ebiten.SetWindowTitle(config.Title)
	ebiten.SetWindowSize(config.Width, config.Height)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game.NewRandom(config.LogicalWidth, config.LogicalHeight)); err != nil {
		log.Fatal(err)
	}
}
