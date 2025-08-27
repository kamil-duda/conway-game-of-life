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

// TODO: tests
func newCanvas(x, y int) *gameCanvas {
	image := ebiten.NewImage(x, y)
	imageOptions := &ebiten.DrawImageOptions{}
	scaleX := float64(config.Width) / float64(image.Bounds().Dx())
	scaleY := float64(config.Height) / float64(image.Bounds().Dy())
	imageOptions.GeoM.Scale(scaleX, scaleY)

	return &gameCanvas{image, imageOptions}
}

// TODO: tests
func (c *gameCanvas) clear() {
	c.image.Clear()
}

// TODO: tests
func (c *gameCanvas) drawOnto(screen *ebiten.Image) {
	screen.DrawImage(c.image, c.options)
}

func (c *gameCanvas) pixel(x, y int) {
	draw.Pixel(x, y, c.image)
}
