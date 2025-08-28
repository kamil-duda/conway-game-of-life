package game

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kamil-duda/conway-game-of-life/config"
)

func TestNewCanvas(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"small canvas", 10, 10},
		{"medium canvas", 100, 100},
		{"large canvas", 1000, 1000},
		{"rectangular canvas", 50, 100},
		{"single draw", 1, 1},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			canvas := newCanvas(tt.width, tt.height)

			if canvas == nil {
				t.Fatal("newCanvas returned nil")
			}
			if canvas.image == nil {
				t.Fatal("canvas image is nil")
			}
			if canvas.options == nil {
				t.Fatal("canvas options is nil")
			}

			bounds := canvas.image.Bounds()
			if bounds.Dx() != tt.width {
				t.Errorf("expected width %d, got %d", tt.width, bounds.Dx())
			}
			if bounds.Dy() != tt.height {
				t.Errorf("expected height %d, got %d", tt.height, bounds.Dy())
			}

			// We can't directly access the scale values from GeoM, but we can verify
			// they were set by checking that GeoM is properly initialized
			if canvas.options.GeoM == (ebiten.GeoM{}) {
				t.Error("GeoM was not initialized with scaling")
			}
		})
	}
}

func BenchmarkNewCanvas(b *testing.B) {
	for b.Loop() {
		_ = newCanvas(100, 100)
	}
}

func BenchmarkCanvasClear(b *testing.B) {
	canvas := newCanvas(100, 100)
	for b.Loop() {
		canvas.clear()
	}
}

func BenchmarkCanvasDraw(b *testing.B) {
	canvas := newCanvas(100, 100)
	for b.Loop() {
		canvas.draw(cell{50, 50})
	}
}

func BenchmarkCanvasDrawOnto(b *testing.B) {
	canvas := newCanvas(100, 100)
	screen := ebiten.NewImage(config.Width, config.Height)
	for i := 0; i < 10; i++ {
		canvas.draw(cell{i, i})
	}

	for b.Loop() {
		canvas.drawOnto(screen)
	}
}
