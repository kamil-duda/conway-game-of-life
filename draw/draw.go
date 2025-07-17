package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func Square(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 10, 10, 10, 10, color.White, true)
}

// todo: configurable - cell size
// todo: configurable - stroke width
// todo: benchmark
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
