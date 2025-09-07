# Conway's Game of Life

## Development

Project uses Makefile for common development tasks.
Available commands:

- `make mod` - Go mod tidy
- `make update` - Update dependencies
- `make run` - Run the application
- `make test` - Run all tests
- `make bench` - Run all benchmarks
- `make coverage` - Generate test coverage report

## Project Management

Project management using GitHub Projects

* [Conway's Game of Life - GitHub Project](https://github.com/users/kamil-duda/projects/2)
* [GitHub Issues](https://github.com/kamil-duda/conway-game-of-life/issues)

## Resources

- [Ebitengine - Cheat Sheet](https://ebitengine.org/en/documents/cheatsheet.html)
- [Ebitengine - Package Documentation](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2)

## To do

1. CountRenderer that can print value with predefined label
    1. Contains *textRenderer
2. PerformanceMonitor
    1. Stores FPS and UPS counters
    2. Used by Game
    3. Contains renderers and is able to render itself onto the screen