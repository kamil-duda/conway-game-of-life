package ui

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

// CreateMainLayout creates the main 2-column layout (game area + right panel)
func CreateMainLayout() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(Theme.Background)),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
				widget.RowLayoutOpts.Spacing(0),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)
}

// CreateMetricsBar creates horizontal bar with 8 metrics in 2 rows (4x2 grid)
func CreateMetricsBar() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(Theme.TopBarBg)),
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(4),
				widget.GridLayoutOpts.Stretch([]bool{true, true, true, true}, []bool{false, false}),
				widget.GridLayoutOpts.Spacing(MetricsSpacing, 0),
				widget.GridLayoutOpts.Padding(&widget.Insets{Top: 4, Right: 8, Bottom: 4, Left: 8}), // 8px horizontal, 4px vertical
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(0, TopBarHeight),
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)
}

// CreateGamePanel creates the left side with metrics bar + game area
func CreateGamePanel(metricsBar *widget.Container) *widget.Container {
	gamePanel := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)

	// Add metrics bar to the top
	gamePanel.AddChild(metricsBar)

	// Add spacer for game area (the actual game canvas will be drawn by ebiten)
	gameArea := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(Theme.Background)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)
	gamePanel.AddChild(gameArea)

	return gamePanel
}

// CreateGameAreaOnly creates just the game area container without metrics bar
func CreateGameAreaOnly() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(Theme.Background)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: true}),
		),
	)
}

// CreateControlPanel creates the right side vertical container with Steel Blue border
func CreateControlPanel() *widget.Container {
	// Create background with Light Steel Blue color
	backgroundImg := image.NewNineSliceColor(Theme.RightPanelBg)

	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(backgroundImg),
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
				widget.RowLayoutOpts.Spacing(ButtonSpacing),
				widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(ButtonSpacing)),
			),
		),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(RightPanelWidth, 0),
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{Stretch: false}),
		),
	)
}
