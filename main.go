package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/game"
)

func main() {
	ebiten.SetWindowTitle(config.Title)
	ebiten.SetWindowSize(config.InitWindowWidth, config.InitWindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSizeLimits(config.MinWindowWidth, config.MinWindowHeight, config.MaxWindowWidth, config.MaxWindowHeight)
	if err := ebiten.RunGame(game.NewRandomGame(config.LogicalWidth, config.LogicalHeight)); err != nil {
		log.Fatal(err)
	}
}
