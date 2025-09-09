package game

import (
	"iter"
	"maps"
)

type cells = map[uint64]cell

type universe struct {
	liveCells cells
}

func newUniverse(cellz ...cell) *universe {
	u := &universe{make(cells)}
	for _, cell := range cellz {
		u.setLive(cell.x, cell.y)
	}
	return u
}

func (u *universe) isLive(x, y int) bool {
	key := toKey(x, y)
	_, exists := u.liveCells[key]
	return exists
}

func (u *universe) setLive(x, y int) {
	key := toKey(x, y)
	u.liveCells[key] = cell{x, y}
}

func (u *universe) setDead(x, y int) {
	delete(u.liveCells, toKey(x, y))
}

func (u *universe) liveNeighbours(x, y int) (count uint8) {
	// NW
	if u.isLive(x-1, y-1) {
		count++
	}
	// W
	if u.isLive(x-1, y) {
		count++
	}
	// SW
	if u.isLive(x-1, y+1) {
		count++
	}
	// NE
	if u.isLive(x+1, y-1) {
		count++
	}
	// E
	if u.isLive(x+1, y) {
		count++
	}
	// SE
	if u.isLive(x+1, y+1) {
		count++
	}
	// N
	if u.isLive(x, y-1) {
		count++
	}
	// S
	if u.isLive(x, y+1) {
		count++
	}
	return count
}

func toKey(x, y int) uint64 {
	return uint64(uint32(x))<<32 | uint64(uint32(y))
}

func (u *universe) cellsIter() iter.Seq[cell] {
	return maps.Values(u.liveCells)
}

func (u *universe) clone() *universe {
	return &universe{maps.Clone(u.liveCells)}
}
