package game

import (
	"iter"
	"maps"
)

type cells = map[uint64]Cell

type Universe struct {
	liveCells cells
}

func New(cellz ...Cell) *Universe {
	u := &Universe{make(cells)}
	for _, cell := range cellz {
		u.setLive(cell.x, cell.y)
	}
	return u
}

func (universe *Universe) isLive(x, y int) bool {
	key := toKey(x, y)
	_, exists := universe.liveCells[key]
	return exists
}

func (universe *Universe) setLive(x, y int) {
	key := toKey(x, y)
	universe.liveCells[key] = Cell{x, y}
}

func (universe *Universe) setDead(x, y int) {
	delete(universe.liveCells, toKey(x, y))
}

func (universe *Universe) liveNeighbours(x, y int) (count uint8) {
	// NW
	if universe.isLive(x-1, y-1) {
		count++
	}
	// W
	if universe.isLive(x-1, y) {
		count++
	}
	// SW
	if universe.isLive(x-1, y+1) {
		count++
	}
	// NE
	if universe.isLive(x+1, y-1) {
		count++
	}
	// E
	if universe.isLive(x+1, y) {
		count++
	}
	// SE
	if universe.isLive(x+1, y+1) {
		count++
	}
	// N
	if universe.isLive(x, y-1) {
		count++
	}
	// S
	if universe.isLive(x, y+1) {
		count++
	}
	return count
}

func toKey(x, y int) uint64 {
	return uint64(uint32(x))<<32 | uint64(uint32(y))
}

func (universe *Universe) cellsIter() iter.Seq[Cell] {
	return maps.Values(universe.liveCells)
}

func (universe *Universe) clone() *Universe {
	return &Universe{maps.Clone(universe.liveCells)}
}
