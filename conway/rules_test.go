package conway

import "testing"

var liveCellSurvivesTestCases = []struct {
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

func TestLiveCellSurvives(t *testing.T) {
	t.Parallel()
	for _, tt := range liveCellSurvivesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := LiveCellSurvives(tt.neighbours)
			if result != tt.expected {
				t.Errorf("LiveCellSurvives(%d) = %v, want %v", tt.neighbours, result, tt.expected)
			}
		})
	}
}

var deadCellRevivesTestCases = []struct {
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

func TestDeadCellRevives(t *testing.T) {
	t.Parallel()
	for _, tt := range deadCellRevivesTestCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := DeadCellRevives(tt.neighbours)
			if result != tt.expected {
				t.Errorf("DeadCellRevives(%d) = %v, want %v", tt.neighbours, result, tt.expected)
			}
		})
	}
}

func BenchmarkLiveCellSurvives(b *testing.B) {
	for b.Loop() {
		LiveCellSurvives(1)
		LiveCellSurvives(2)
		LiveCellSurvives(3)
		LiveCellSurvives(4)
	}
}

func BenchmarkDeadCellRevives(b *testing.B) {
	for b.Loop() {
		DeadCellRevives(1)
		DeadCellRevives(2)
		DeadCellRevives(3)
		DeadCellRevives(4)
	}
}
