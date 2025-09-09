package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
	"github.com/kamil-duda/conway-game-of-life/draw"
)

type gameCanvas struct {
	image   *ebiten.Image
	options *ebiten.DrawImageOptions
}

func newCanvas(x, y int) *gameCanvas {
	image := ebiten.NewImage(x, y)
	options := &ebiten.DrawImageOptions{}
	scaleX := float64(config.Width) / float64(image.Bounds().Dx())
	scaleY := float64(config.Height) / float64(image.Bounds().Dy())
	options.GeoM.Scale(scaleX, scaleY)
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
