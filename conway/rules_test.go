package conway

import "testing"

var checkLiveTestCases = []struct {
	name       string
	neighbours uint8
	expected   bool
}{
	{"0 neighbours - dies", 0, false},
	{"1 neighbour - dies", 1, false},
	{"2 neighbours - survives", 2, true},
	{"3 neighbours - survives", 3, true},
	{"4 neighbours - dies", 4, false},
	{"5 neighbours - dies", 5, false},
	{"6 neighbours - dies", 6, false},
	{"7 neighbours - dies", 7, false},
	{"8 neighbours - dies", 8, false},
}

func TestCheckLiveCell(t *testing.T) {
	t.Parallel()
	for _, tt := range checkLiveTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := checkLiveCell(tt.neighbours)
			if result != tt.expected {
				t.Errorf("checkLiveCell(%d) = %v, want %v", tt.neighbours, result, tt.expected)
			}
		})
	}
}

var checkDeadTestCases = []struct {
	name       string
	neighbours uint8
	expected   bool
}{
	{"0 neighbours - stays dead", 0, false},
	{"1 neighbour - stays dead", 1, false},
	{"2 neighbours - stays dead", 2, false},
	{"3 neighbours - becomes alive", 3, true},
	{"4 neighbours - stays dead", 4, false},
	{"5 neighbours - stays dead", 5, false},
	{"6 neighbours - stays dead", 6, false},
	{"7 neighbours - stays dead", 7, false},
	{"8 neighbours - stays dead", 8, false},
}

func TestCheckDeadCell(t *testing.T) {
	t.Parallel()
	for _, tt := range checkDeadTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := checkDeadCell(tt.neighbours)
			if result != tt.expected {
				t.Errorf("checkDeadCell(%d) = %v, want %v", tt.neighbours, result, tt.expected)
			}
		})
	}
}

func BenchmarkCheckLiveCell(b *testing.B) {
	for b.Loop() {
		checkLiveCell(1)
		checkLiveCell(2)
		checkLiveCell(3)
		checkLiveCell(4)
	}
}

func BenchmarkCheckDeadCell(b *testing.B) {
	for b.Loop() {
		checkDeadCell(1)
		checkDeadCell(2)
		checkDeadCell(3)
		checkDeadCell(4)
	}
}
