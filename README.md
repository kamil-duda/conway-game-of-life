# Conway's Game of Life

## Resources

- [Ebitengine - Cheat Sheet](https://ebitengine.org/en/documents/cheatsheet.html)
- [Ebitengine - Package Documentation](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2)

## To do

### Core

- [x] Project structure setup - go modules, basic files
- [x] Draw window with square - basic Ebiten window setup
- [ ] Divide background into cells - grid system
- [ ] Cell-to-pixel coordinate translation system
- [ ] Basic cell rendering - draw alive/dead cells
- [ ] Screen configuration from external config - YAML config file
- [ ] Game state data structure - 2D grid/array for cells
- [ ] Game of Life rules implementation - birth, survival, death rules
- [ ] Game logic working on cells not pixels - abstract cell operations
- [ ] Update function to update state - game loop logic

### User Interface and Interaction

- [ ] Play and pause functionality
- [ ] Clicking on cells toggles live/dead state
- [ ] Mouse drag to toggle multiple cells
- [ ] Reset button
- [ ] Clear grid button
- [ ] FPS
- [ ] Generation counter display
- [ ] Population counter display
- [ ] Slider for speed control
- [ ] Toolbox with pre-made patterns - glider, blinker, etc.
- [ ] Pattern placement system
- [ ] Save/load grid states to a file
- [ ] Grid lines toggle

### Improvements

- [ ] Resizing working - dynamic window resizing
- [ ] Zoom in/out functionality
- [ ] Pan/scroll the grid
- [ ] Performance optimization - only update changed regions
- [ ] Shader-based rendering for better performance

### Code Quality

- [ ] Unit tests for game logic
- [ ] Integration tests for UI interactions
- [ ] Performance benchmarks in Go
- [ ] Code documentation

### CI/CD & Build

- [ ] GitHub Actions workflow for tests, build, benchmarks
- [ ] Cross-platform build - windows, macOS, Linux
- [ ] Release automation with GitHub releases