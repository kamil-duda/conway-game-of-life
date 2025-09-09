package draw

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gomono"
)

type textRenderer struct {
	font *text.GoTextFace
	opts *text.DrawOptions
}

func newTextRenderer(fontSize float64, x, y float64) *textRenderer {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		log.Fatal("Failed to initialize fonts:", err)
	}
	font := &text.GoTextFace{Source: fontSource, Size: fontSize}

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(x, y)

	return &textRenderer{font, opts}
}

func (r *textRenderer) draw(contents string, screen *ebiten.Image) {
	text.Draw(screen, contents, r.font, r.opts)
}
