package draw

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func DrawSquare(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 10, 10, 10, 10, color.White, true)
}
