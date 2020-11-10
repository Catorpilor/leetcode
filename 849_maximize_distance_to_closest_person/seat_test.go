package seat

import "testing"

func TestMaxDistToClosest(t *testing.T) {
	st := []struct {
		name  string
		seats []int
		exp   int
	}{
		{"only one seat left", []int{1, 1, 0}, 1},
		{"testcase1", []int{1, 0, 0, 0}, 3},
		{"testcase2", []int{1, 0, 0, 1, 0, 1}, 1},
		{"testcase3", []int{1, 0, 0, 0, 1, 0, 1}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDist(tt.seats)
			if out != tt.exp {
				t.Fatalf("with input seats:%v wanted %d but got %d", tt.seats, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
