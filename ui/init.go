package ui

import (
	"github.com/ebitenui/ebitenui/widget"
)

// MetricsWidgets holds references to the metric labels for updating
type MetricsWidgets struct {
	FpsLabel      *widget.Label
	UpsLabel      *widget.Label
	GenLabel      *widget.Label
	PopLabel      *widget.Label
	AliveLabel    *widget.Label
	DeadLabel     *widget.Label
	ActivityLabel *widget.Label
	RuntimeLabel  *widget.Label
}

// InitializeUI creates and initializes the complete UI layout
func InitializeUI(resetCallback func(), pauseCallback func(), stepCallback func(), patternCallback func(), settingsCallback func()) (*widget.Container, *MetricsWidgets) {
	// Create the root container with vertical layout
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewRowLayout(
				widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			),
		),
	)

	// Create metrics bar with labels for two rows (4x2 grid)
	metricsBar := CreateMetricsBar()
	// Row 1 labels
	fpsLabel := CreateMetricLabel("🎯 FPS: 0")
	upsLabel := CreateMetricLabel("⚡ UPS: 0")
	genLabel := CreateMetricLabel("🧬 Generation: 0")
	popLabel := CreateMetricLabel("👥 Population: 0")
	// Row 2 labels
	aliveLabel := CreateMetricLabel("📊 Alive: 0")
	deadLabel := CreateMetricLabel("💀 Dead: 0")
	activityLabel := CreateMetricLabel("🔄 Activity: Low")
	runtimeLabel := CreateMetricLabel("⏱️ Runtime: 0:00")

	// Add metric labels to metrics bar (row by row)
	metricsBar.AddChild(fpsLabel)
	metricsBar.AddChild(upsLabel)
	metricsBar.AddChild(genLabel)
	metricsBar.AddChild(popLabel)
	metricsBar.AddChild(aliveLabel)
	metricsBar.AddChild(deadLabel)
	metricsBar.AddChild(activityLabel)
	metricsBar.AddChild(runtimeLabel)

	// Add metrics bar to root (spans full width)
	rootContainer.AddChild(metricsBar)

	// Create horizontal container for the 2-column layout
	bottomLayout := CreateMainLayout()

	// Create game area (left side) - without metrics bar
	gameArea := CreateGameAreaOnly()

	// Create control panel (right side)
	controlPanel := CreateControlPanel()

	// Create control buttons
	resetButton := CreateStyledButton("RESET", func(args *widget.ButtonClickedEventArgs) {
		resetCallback()
	})

	pauseButton := CreateStyledButton("PLAY/PAUSE", func(args *widget.ButtonClickedEventArgs) {
		pauseCallback()
	})

	stepButton := CreateStyledButton("SINGLE STEP", func(args *widget.ButtonClickedEventArgs) {
		stepCallback()
	})

	patternButton := CreateStyledButton("PATTERN LOADER", func(args *widget.ButtonClickedEventArgs) {
		patternCallback()
	})

	settingsButton := CreateStyledButton("SETTINGS", func(args *widget.ButtonClickedEventArgs) {
		settingsCallback()
	})

	// Add buttons to control panel in order specified by refactor plan
	controlPanel.AddChild(resetButton)
	controlPanel.AddChild(pauseButton)
	controlPanel.AddChild(stepButton)
	controlPanel.AddChild(patternButton)
	controlPanel.AddChild(settingsButton)

	// Add panels to bottom layout
	bottomLayout.AddChild(gameArea)
	bottomLayout.AddChild(controlPanel)

	// Add bottom layout to root
	rootContainer.AddChild(bottomLayout)

	// Return the root container and metrics widgets
	metricsWidgets := &MetricsWidgets{
		FpsLabel:      fpsLabel,
		UpsLabel:      upsLabel,
		GenLabel:      genLabel,
		PopLabel:      popLabel,
		AliveLabel:    aliveLabel,
		DeadLabel:     deadLabel,
		ActivityLabel: activityLabel,
		RuntimeLabel:  runtimeLabel,
	}

	return rootContainer, metricsWidgets
}
