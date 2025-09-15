package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/draw"
	"github.com/kamil-duda/conway-game-of-life/ui"
)

type gameCanvas struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func newCanvas(x, y int) *gameCanvas {
	image := ebiten.NewImage(x, y)
	options := &ebiten.DrawImageOptions{}
	// Calculate available game area (excluding UI elements)
	availableWidth := config.InitWindowWidth - ui.RightPanelWidth
	availableHeight := config.InitWindowHeight - ui.TopBarHeight
	scaleX := float64(availableWidth) / float64(image.Bounds().Dx())
	scaleY := float64(availableHeight) / float64(image.Bounds().Dy())
	options.GeoM.Scale(scaleX, scaleY)
	// Offset to account for the top bar
	options.GeoM.Translate(0, float64(ui.TopBarHeight))
	return &gameCanvas{image, options}
}

func (gc *gameCanvas) clear() {
	gc.image.Clear()
}

func (gc *gameCanvas) drawOnto(screen *ebiten.Image) {
	screen.DrawImage(gc.image, gc.options)
}

func (gc *gameCanvas) draw(c cell) {
	draw.Pixel(c.x, c.y, gc.image)
}
