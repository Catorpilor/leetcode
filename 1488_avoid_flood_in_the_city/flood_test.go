package flood

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAvoidFlood(t *testing.T) {
	st := []struct {
		name string
		rain []int
		exp  []int
	}{
		{"one rainy day", []int{1}, []int{-1}},
		{"one sunny day", []int{0}, []int{1}},
		{"all rainy days", []int{1, 2, 3, 4}, []int{-1, -1, -1, -1}},
		{"flood city", []int{1, 1, 0}, []int{}},
		{"testcase1", []int{1, 2, 0, 0, 2, 1}, []int{-1, -1, 2, 1, -1, -1}},
		{"failed1", []int{0, 72328, 0, 0, 94598, 54189, 39171, 53361, 0, 0, 0, 72742, 0, 98613, 16696, 0, 32756, 23537, 0, 94598, 0, 0, 0, 11594, 27703, 0, 0, 0, 20081, 0, 24645, 0, 0, 0, 0, 0, 0, 0, 2711, 98613, 0, 0, 0, 0, 0, 91987, 0, 0, 0, 22762, 23537, 0, 0, 0, 0, 54189, 0, 0, 87770, 0, 0, 0, 0, 27703, 0, 0, 0, 0, 20081, 16696, 0, 0, 0, 0, 0, 0, 0, 35903, 0, 72742, 0, 0, 0, 35903, 0, 0, 91987, 76728, 0, 0, 0, 0, 2711, 0, 0, 11594, 0, 0, 22762, 24645, 0, 0, 0, 0, 0, 53361, 0, 87770, 0, 0, 39171, 0, 24577, 0, 0, 0, 24577, 0, 0, 72328, 0, 0, 32756, 0, 0, 76728}, []int{1, -1, 72328, 1, -1, -1, -1, -1, 94598, 54189, 53361, -1, 72742, -1, -1, 98613, -1, -1, 23537, -1, 16696, 39171, 32756, -1, -1, 27703, 11594, 1, -1, 20081, -1, 24645, 1, 1, 1, 1, 1, 1, -1, -1, 2711, 1, 1, 1, 1, -1, 91987, 1, 1, -1, -1, 22762, 1, 1, 1, -1, 1, 1, -1, 87770, 1, 1, 1, -1, 1, 1, 1, 1, -1, -1, 1, 1, 1, 1, 1, 1, 1, -1, 35903, -1, 1, 1, 1, -1, 1, 1, -1, -1, 76728, 1, 1, 1, -1, 1, 1, -1, 1, 1, -1, -1, 1, 1, 1, 1, 1, -1, 1, -1, 1, 1, -1, 1, -1, 24577, 1, 1, -1, 1, 1, -1, 1, 1, -1, 1, 1, -1}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := avoidFlood(tt.rain)
			if diff := cmp.Diff(out, tt.exp); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
			t.Log("pass")
		})
	}
}

func TestLowserBound(t *testing.T) {
	st := []struct {
		name        string
		arr         []int
		target, exp int
	}{
		{"all identical", []int{1, 1, 1}, 2, -1},
		{"testcase1", []int{1, 2, 4, 5}, 3, 2},
		{"testcase2", []int{1, 2, 3, 4, 5}, 3, 3},
		{"testcase3", []int{2, 3}, 1, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := lowerBound(tt.arr, tt.target)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and target:%d wanted %d but got %d", tt.arr, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
