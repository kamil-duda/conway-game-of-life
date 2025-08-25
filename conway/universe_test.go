package conway

import (
	"math"
	"reflect"
	"testing"
)

func TestNewUniverseNotNil(t *testing.T) {
	t.Parallel()
	got := New()
	if got == nil {
		t.Error("New() returned nil, want non-nil Universe")
	}
}

func TestNewUniverseEmpty(t *testing.T) {
	t.Parallel()
	got := len(New().liveCells)
	want := 0
	if got != want {
		t.Errorf("Got some live cells %v, but want %v", got, want)
	}
}

var testToKeyCases = []struct {
	name string
	x, y int
	want uint64
}{
	{"(0, 0) is 0", 0, 0, 0},
	{"(0, -0) is 0", 0, -0, 0},
	{"(-0, 0) is 0", -0, 0, 0},
	{"(-0, -0) is 0", -0, -0, 0},
	{"(0, 1) is 1", 0, 1, 1},
	{"(1, 0) is 0x0000000100000000", 1, 0, 0x0000000100000000},
	{"(1, 1) is 0x0000000100000001", 1, 1, 0x0000000100000001},
	{"(0, -1) is 0x00000000FFFFFFFF", 0, -1, 0x00000000FFFFFFFF},
	{"(-1, 0) is 0xFFFFFFFF00000000", -1, 0, 0xFFFFFFFF00000000},
	{"(-1, -1) is 0xFFFFFFFFFFFFFFFF", -1, -1, 0xFFFFFFFFFFFFFFFF},
	{"(max, max) is 0x7FFFFFFF7FFFFFFF", math.MaxInt32, math.MaxInt32, 0x7FFFFFFF7FFFFFFF},
	{"(min, min) is 0x1000000010000000", math.MinInt32, math.MinInt32, 0x8000000080000000},
}

func TestToKey(t *testing.T) {
	t.Parallel()
	for _, tt := range testToKeyCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := toKey(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var isLiveTestCases = []struct {
	name      string
	liveCells grid
	query     uint64
	want      bool
}{
	{"empty universe", make(grid), toKey(0, 0), false},
	{"non-empty universe, live", map[uint64]bool{toKey(0, 0): true}, toKey(0, 0), true},
	{"non-empty universe, dead", map[uint64]bool{toKey(1, 1): true}, toKey(0, 0), false},
}

func TestIsLive(t *testing.T) {
	t.Parallel()
	for _, tt := range isLiveTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			universe := Universe{tt.liveCells}
			got := universe.isLive(0, 0)
			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setLiveTestCases = []struct {
	name      string
	liveCells grid
	x, y      int
	want      grid
}{
	{"empty universe", make(grid), 0, 0, map[uint64]bool{toKey(0, 0): true}},
	{"non-empty universe, same cell", map[uint64]bool{toKey(0, 0): true}, 0, 0, map[uint64]bool{toKey(0, 0): true}},
	{"non-empty universe, different cell", map[uint64]bool{toKey(0, 0): true}, 1, 1, map[uint64]bool{toKey(0, 0): true, toKey(1, 1): true}},
}

func TestSetLive(t *testing.T) {
	t.Parallel()
	for _, tt := range setLiveTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			universe := Universe{tt.liveCells}
			universe.setLive(tt.x, tt.y)
			got := universe.liveCells
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setDeadTestCases = []struct {
	name      string
	liveCells grid
	x, y      int
	want      grid
}{
	{"empty universe", make(grid), 0, 0, make(grid)},
	{"non-empty universe, same cell", map[uint64]bool{toKey(0, 0): true}, 0, 0, make(grid)},
	{"non-empty universe, different cell", map[uint64]bool{toKey(0, 0): true}, 1, 1, map[uint64]bool{toKey(0, 0): true}},
}

func TestSetDead(t *testing.T) {
	t.Parallel()
	for _, tt := range setDeadTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			universe := Universe{tt.liveCells}
			universe.setDead(tt.x, tt.y)
			got := universe.liveCells
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setLiveNeighboursCases = []struct {
	name      string
	liveCells grid
	x, y      int
	want      uint8
}{
	{"empty universe", make(grid), 0, 0, 0},
	{"one neighbour", map[uint64]bool{toKey(-1, 0): true}, 0, 0, 1},
	{"two neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
	}, 0, 0, 2},
	{"three neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
	}, 0, 0, 3},
	{"four neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
	}, 0, 0, 4},
	{"five neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
		toKey(0, -1):  true,
	}, 0, 0, 5},
	{"six neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
		toKey(0, -1):  true,
		toKey(1, 0):   true,
	}, 0, 0, 6},
	{"seven neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
		toKey(0, -1):  true,
		toKey(1, 0):   true,
		toKey(1, 1):   true},
		0, 0, 7},
	{"eight neighbours", map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
		toKey(0, -1):  true,
		toKey(1, 0):   true,
		toKey(1, 1):   true,
		toKey(1, -1):  true},
		0, 0, 8},
}

func TestLiveNeighbours(t *testing.T) {
	t.Parallel()
	for _, tt := range setLiveNeighboursCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			universe := Universe{tt.liveCells}
			got := universe.liveNeighbours(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewUniverse(b *testing.B) {
	for b.Loop() {
		New()
	}
}

func BenchmarkToKey(b *testing.B) {
	for b.Loop() {
		toKey(0, 0)
	}
}

func BenchmarkIsLive(b *testing.B) {
	u := Universe{map[uint64]bool{toKey(0, 0): true}}
	for b.Loop() {
		u.isLive(0, 0)
	}
}

func BenchmarkSetLive(b *testing.B) {
	u := Universe{map[uint64]bool{toKey(0, 0): true}}
	for b.Loop() {
		u.setLive(1, 1)
	}
}

func BenchmarkLiveNeighbours(b *testing.B) {
	u := Universe{map[uint64]bool{
		toKey(-1, -1): true,
		toKey(-1, 0):  true,
		toKey(-1, 1):  true,
		toKey(0, 1):   true,
		toKey(0, -1):  true,
		toKey(1, 0):   true,
		toKey(1, 1):   true,
		toKey(1, -1):  true}}
	for b.Loop() {
		u.liveNeighbours(0, 0)
	}
}

func BenchmarkSetDead(b *testing.B) {
	u := Universe{map[uint64]bool{toKey(0, 0): true}}
	for b.Loop() {
		u.setDead(0, 0)
	}
}
