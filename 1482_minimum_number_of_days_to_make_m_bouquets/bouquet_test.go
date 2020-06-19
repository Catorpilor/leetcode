package bouquet

import "testing"

func TestMinDays(t *testing.T) {
	st := []struct {
		name      string
		bloomDay  []int
		m, k, exp int
	}{
		{"testcase1", []int{1, 10, 3, 10, 2}, 3, 1, 3},
		{"testcase2", []int{1, 10, 3, 10, 2}, 3, 2, -1},
		{"testcase3", []int{7, 7, 7, 7, 12, 7, 7}, 2, 3, 12},
		{"testcase4", []int{1000000000, 1000000000}, 1, 1, 1000000000},
		{"testcase5", []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2, 9},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minDays(tt.bloomDay, tt.m, tt.k)
			if out != tt.exp {
				t.Fatalf("with input bloomDay:%v m:%d, k:%d wanted %d but got %d", tt.bloomDay, tt.m, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
