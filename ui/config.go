package ui

import (
	"image/color"
)

const (
	TopBarHeight    = 60
	RightPanelWidth = 150
	MetricsSpacing  = 20
	ButtonSpacing   = 10
	MetricsFontSize = 12
	ButtonFontSize  = 11
)

// Color definitions based on UI refactor plan
var (
	SteelBlue      = color.RGBA{70, 130, 180, 255}  // #4682B4
	LightSteelBlue = color.RGBA{176, 196, 222, 255} // #B0C4DE
	LightCyan      = color.RGBA{224, 255, 255, 255} // #E0FFFF
	White          = color.RGBA{255, 255, 255, 255} // #FFFFFF
	LightGray      = color.RGBA{211, 211, 211, 255} // #D3D3D3
	DarkGray       = color.RGBA{47, 47, 47, 255}    // #2F2F2F
	Black          = color.RGBA{0, 0, 0, 255}       // #000000
)

var Theme = struct {
	Background       color.Color
	TopBarBg         color.Color
	RightPanelBg     color.Color
	RightPanelBorder color.Color
	MetricsText      color.Color
	ButtonIdle       color.Color
	ButtonHover      color.Color
	ButtonPressed    color.Color
	ButtonText       color.Color
	ButtonTextHover  color.Color
	GridLines        color.Color
	LiveCell         color.Color
	DeadCell         color.Color
}{
	Background:       LightCyan,
	TopBarBg:         SteelBlue,
	RightPanelBg:     LightSteelBlue,
	RightPanelBorder: SteelBlue,
	MetricsText:      White,
	ButtonIdle:       White,
	ButtonHover:      LightGray,
	ButtonPressed:    SteelBlue,
	ButtonText:       DarkGray,
	ButtonTextHover:  DarkGray,
	GridLines:        LightGray,
	LiveCell:         Black,
	DeadCell:         White,
}
