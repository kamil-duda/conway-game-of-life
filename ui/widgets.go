package ui

import (
	"bytes"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gomono"
)

// CreateMetricLabel creates styled labels for metrics
func CreateMetricLabel(text string) *widget.Label {
	font := DefaultFont()
	return widget.NewLabel(
		widget.LabelOpts.Text(text, &font, &widget.LabelColor{
			Idle:     Theme.MetricsText,
			Disabled: Theme.MetricsText,
		}),
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
			}),
		)),
	)
}

// CreateStyledButton creates nice buttons with hover/pressed states
func CreateStyledButton(text string, clickHandler widget.ButtonClickedHandlerFunc) *widget.Button {
	font := ButtonFont()

	// Create button images (border radius would need custom implementation)
	idleImg := image.NewNineSliceColor(Theme.ButtonIdle)
	hoverImg := image.NewNineSliceColor(Theme.ButtonHover)
	pressedImg := image.NewNineSliceColor(Theme.ButtonPressed)

	return widget.NewButton(
		widget.ButtonOpts.TextLabel(text),
		widget.ButtonOpts.TextFace(&font),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:     Theme.ButtonText,
			Disabled: Theme.ButtonText,
			Hover:    Theme.ButtonTextHover,
			Pressed:  White, // White text on steel blue pressed state
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:         idleImg,
			Hover:        hoverImg,
			Pressed:      pressedImg,
			PressedHover: pressedImg,
			Disabled:     idleImg,
		}),
		widget.ButtonOpts.ClickedHandler(clickHandler),
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(120, 35), // Updated to match plan: 120px×35px
		),
	)
}

// DefaultFont returns the monospace font for metrics (as per refactor plan)
func DefaultFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   MetricsFontSize,
	}
}

// ButtonFont returns the bold sans-serif font for buttons (as per refactor plan)
func ButtonFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(gobold.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   ButtonFontSize,
	}
}
