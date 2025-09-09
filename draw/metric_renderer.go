package draw

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type MetricRenderer struct {
	*textRenderer
	label string
}

func (r *MetricRenderer) Draw(value int, screen *ebiten.Image) {
	r.draw(fmt.Sprintf("%s%d", r.label, value), screen)
}

func NewFpsRenderer() *MetricRenderer {
	return &MetricRenderer{newTextRenderer(24, 0, 0), "FPS: "}
}

func NewSimSpeedRenderer() *MetricRenderer {
	return &MetricRenderer{newTextRenderer(24, 200, 0), "UPS: "}
}

func NewGenerationRenderer() *MetricRenderer {
	return &MetricRenderer{newTextRenderer(24, 400, 0), "Gen: "}
}

func NewPopulationRenderer() *MetricRenderer {
	return &MetricRenderer{newTextRenderer(24, 600, 0), "Pop: "}
}
