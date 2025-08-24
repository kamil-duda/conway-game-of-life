package conway

type grid = map[uint64]bool

type Universe struct {
	liveCells grid
}

func New() *Universe {
	return new(Universe)
}

func (universe *Universe) IsLive(x, y int) bool {
	key := toKey(x, y)
	_, exists := universe.liveCells[key]
	return exists
}

func (universe *Universe) SetLive(x, y int) {
	key := toKey(x, y)
	universe.liveCells[key] = true
}

func toKey(x, y int) uint64 {
	return uint64(uint32(x))<<32 | uint64(uint32(y))
}
