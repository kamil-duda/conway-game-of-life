package game

import (
	"math"
	"reflect"
	"slices"
	"testing"
)

func TestNewUniverseNotNil(t *testing.T) {
	t.Parallel()
	got := newUniverse()
	if got == nil {
		t.Error("newUniverse() returned nil, want non-nil universe")
	}
}

func TestNewUniverseEmpty(t *testing.T) {
	t.Parallel()
	got := len(newUniverse().liveCells)
	want := 0
	if got != want {
		t.Errorf("Got some live cells %v, but want %v", got, want)
	}
}

func TestNewUniverseSetLive(t *testing.T) {
	t.Parallel()
	u := newUniverse()
	u.setLive(0, 0)
	got := len(u.liveCells)
	want := 1
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
	name     string
	universe *universe
	query    uint64
	want     bool
}{
	{"empty universe", newUniverse(), toKey(0, 0), false},
	{"non-empty universe, live", newUniverse(cell{0, 0}), toKey(0, 0), true},
	{"non-empty universe, dead", newUniverse(cell{1, 1}), toKey(0, 0), false},
}

func TestIsLive(t *testing.T) {
	t.Parallel()
	for _, tt := range isLiveTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.universe.isLive(0, 0)
			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setLiveTestCases = []struct {
	name     string
	universe *universe
	x, y     int
	want     cells
}{
	{"empty universe", newUniverse(), 0, 0, newUniverse(cell{0, 0}).liveCells},
	{"non-empty universe, same cell", newUniverse(cell{0, 0}), 0, 0, newUniverse(cell{0, 0}).liveCells},
	{"non-empty universe, different cell", newUniverse(cell{0, 0}), 1, 1, newUniverse(cell{0, 0}, cell{1, 1}).liveCells},
}

func TestSetLive(t *testing.T) {
	t.Parallel()
	for _, tt := range setLiveTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.universe.setLive(tt.x, tt.y)
			got := tt.universe.liveCells
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setDeadTestCases = []struct {
	name     string
	universe *universe
	x, y     int
	want     cells
}{
	{"empty universe", newUniverse(), 0, 0, make(cells)},
	{"non-empty universe, same cell", newUniverse(cell{0, 0}), 0, 0, make(cells)},
	{"non-empty universe, different cell", newUniverse(cell{0, 0}), 1, 1, newUniverse(cell{0, 0}).liveCells},
}

func TestSetDead(t *testing.T) {
	t.Parallel()
	for _, tt := range setDeadTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.universe.setDead(tt.x, tt.y)
			got := tt.universe.liveCells
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

var setLiveNeighboursCases = []struct {
	name     string
	universe *universe
	x, y     int
	want     uint8
}{
	{"empty universe", newUniverse(), 0, 0, 0},
	{"one neighbour", newUniverse(cell{-1, 0}), 0, 0, 1},
	{"two neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
	), 0, 0, 2},
	{"three neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
	), 0, 0, 3},
	{"four neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
	), 0, 0, 4},
	{"five neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
	), 0, 0, 5},
	{"six neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
	), 0, 0, 6},
	{"seven neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
		cell{1, 1},
	), 0, 0, 7},
	{"eight neighbours", newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
		cell{1, 1},
		cell{1, -1},
	), 0, 0, 8},
}

func TestLiveNeighbours(t *testing.T) {
	t.Parallel()
	for _, tt := range setLiveNeighboursCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.universe.liveNeighbours(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCellsIter(t *testing.T) {
	var testCases = []struct {
		name  string
		cells []cell
	}{
		{"universe contains cells", []cell{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, 1},
			{0, -1},
			{1, 0},
			{1, 1},
			{1, -1},
		}},
		{"universe is empty", []cell{}},
	}

	t.Parallel()
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := newUniverse(tt.cells...)
			got := slices.Collect(u.cellsIter())
			gotSet := make(map[cell]bool)
			for _, cell := range got {
				gotSet[cell] = true
			}
			wantSet := make(map[cell]bool)
			for _, cell := range tt.cells {
				wantSet[cell] = true
			}

			if !reflect.DeepEqual(gotSet, wantSet) {
				t.Errorf("Got %v, want %v", got, tt.cells)
			}
			if len(got) != len(tt.cells) {
				t.Errorf("Got %d cells, want %d", len(got), len(tt.cells))
			}
		})
	}
}

func TestClone(t *testing.T) {
	t.Parallel()

	cells := []cell{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	original := newUniverse(cells...)
	cloned := original.clone()

	if original == cloned {
		t.Errorf("clone() returned same universe pointer %p, want different pointer", original)
	}

	if &original.liveCells == &cloned.liveCells {
		t.Errorf("clone() shares same liveCells map pointer %p, want different pointer", &original.liveCells)
	}

	for cell := range original.cellsIter() {
		if !cloned.isLive(cell.x, cell.y) {
			t.Errorf("cloned universe missing cell (%d, %d)", cell.x, cell.y)
		}
	}
}

func BenchmarkNewUniverse(b *testing.B) {
	for b.Loop() {
		newUniverse()
	}
}

func BenchmarkToKey(b *testing.B) {
	for b.Loop() {
		toKey(0, 0)
	}
}

func BenchmarkIsLive(b *testing.B) {
	u := newUniverse(cell{0, 0})
	for b.Loop() {
		u.isLive(0, 0)
	}
}

func BenchmarkSetLive(b *testing.B) {
	u := newUniverse(cell{0, 0})
	for b.Loop() {
		u.setLive(1, 1)
	}
}

func BenchmarkLiveNeighbours(b *testing.B) {
	u := newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
		cell{1, 1},
		cell{1, -1},
	)
	for b.Loop() {
		u.liveNeighbours(0, 0)
	}
}

func BenchmarkSetDead(b *testing.B) {
	u := newUniverse(cell{0, 0})
	for b.Loop() {
		u.setDead(0, 0)
	}
}

func BenchmarkCellsIter(b *testing.B) {
	u := newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
		cell{1, 1},
		cell{1, -1},
	)
	for b.Loop() {
		u.cellsIter()
	}
}

func BenchmarkClone(b *testing.B) {
	u := newUniverse(
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
		cell{0, 1},
		cell{0, -1},
		cell{1, 0},
		cell{1, 1},
		cell{1, -1},
	)
	for b.Loop() {
		u.clone()
	}
}
