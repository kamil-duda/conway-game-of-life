package draw

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type FpsRenderer struct {
	*textRenderer
}

func NewFpsRenderer() *FpsRenderer {
	return &FpsRenderer{newTextRenderer(24, 0, 0)}
}

func (r *FpsRenderer) Draw(value uint, screen *ebiten.Image) {
	r.draw(fmt.Sprintf("FPS: %d", value), screen)
}
