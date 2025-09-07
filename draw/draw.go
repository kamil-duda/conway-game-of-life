package draw

import (
	"bytes"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/gomono"
)

func Pixel(x int, y int, screen *ebiten.Image) {
	screen.Set(x, y, color.White)
}

func Fps(value uint, screen *ebiten.Image) {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		panic(err)
	}
	face := &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
	op := &text.DrawOptions{}
	op.GeoM.Translate(0, 0)
	text.Draw(screen, fmt.Sprintf("FPS: %v", value), face, op)
}

func DebugBackground(screen *ebiten.Image) {
	vector.StrokeLine(screen, 1, -200, 1, 200, 1, color.RGBA{R: 255, A: 64}, false)
	vector.StrokeLine(screen, -200, 1, 200, 1, 1, color.RGBA{R: 255, A: 64}, false)
}

func Background(screen *ebiten.Image) {
	cellSize := 50
	strokeWidth := float32(2)
	strokeWidthOffset := strokeWidth / 2
	screenWidth := float32(screen.Bounds().Dx())
	screenHeight := float32(screen.Bounds().Dy())

	for verticalIndex := 0; ; verticalIndex++ {
		x := float32(cellSize * verticalIndex)
		adjustedX := x + strokeWidthOffset
		if x > screenWidth {
			break
		}
		vector.StrokeLine(screen, adjustedX, 0, adjustedX, screenHeight, strokeWidth, color.White, false)
	}

	for horizontalIndex := 0; ; horizontalIndex++ {
		y := float32(cellSize * horizontalIndex)
		adjustedY := y + strokeWidthOffset
		if y > screenHeight {
			break
		}
		vector.StrokeLine(screen, 0, adjustedY, screenWidth, adjustedY, strokeWidth, color.White, false)
	}
}
