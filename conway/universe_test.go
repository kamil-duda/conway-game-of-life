package conway

import (
	"reflect"
	"testing"
)

func TestNewUniverseNotNil(t *testing.T) {
	got := New()

	if got == nil {
		t.Error("New() returned nil, want non-nil Universe")
	}
}

func TestNewUniverseEmptyGrid(t *testing.T) {
	got := len(New().grid)
	want := 0

	if got != want {
		t.Errorf("Got Grid of length %v, but want %v", got, want)
	}
}

var universeTestCases = []struct {
	name    string
	initial Grid
	want    Grid
}{
	{name: "underpopulation - 0 neighbours",
		initial: Grid{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
		want: Grid{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		},
	},
	{name: "underpopulation - 1 neighbour",
		initial: Grid{
			{false, false, false},
			{false, true, true},
			{false, false, false},
		},
		want: Grid{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		},
	},
	{name: "survive - 2 neighbours - horizontal",
		initial: Grid{
			{false, false, false},
			{true, true, true},
			{false, false, false},
		},
		want: Grid{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	},
	{name: "survive - 2 neighbours - vertical",
		initial: Grid{
			{false, true, false},
			{false, true, false},
			{false, true, false},
		},
		want: Grid{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	},
	{name: "survive - 2 neighbours - diagonal A",
		initial: Grid{
			{true, false, false},
			{false, true, false},
			{false, false, true},
		},
		want: Grid{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	},
	{name: "survive - 2 neighbours - diagonal B",
		initial: Grid{
			{false, false, true},
			{false, true, false},
			{true, false, false},
		},
		want: Grid{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	},
}

func TestUniverseRules(t *testing.T) {
	t.Parallel()
	for _, test := range universeTestCases {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			universe := Universe{test.initial}
			universe.Update()
			got := universe.grid
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Got %v, want %v", got, test.want)
			}
		})
	}

}

func BenchmarkUniverse(b *testing.B) {
	for b.Loop() {
		New()
	}
}
