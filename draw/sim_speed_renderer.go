package draw

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type SimSpeedRenderer struct {
	*textRenderer
}

func NewSimSpeedRenderer() *SimSpeedRenderer {
	return &SimSpeedRenderer{newTextRenderer(24, 200, 0)}
}

func (r *SimSpeedRenderer) Draw(value uint, screen *ebiten.Image) {
	r.draw(fmt.Sprintf("UPS: %d", value), screen)
}
