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
func (canvas *gameCanvas) clear() {
	canvas.image.Clear()
}

// TODO: tests
func (canvas *gameCanvas) drawOnto(screen *ebiten.Image) {
	screen.DrawImage(canvas.image, canvas.options)
}

func (canvas *gameCanvas) pixel(x, y int) {
	draw.Pixel(x, y, canvas.image)
}
