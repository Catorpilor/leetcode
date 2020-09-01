package pattern

import "testing"

func TestHavePattern(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		m, k int
		exp  bool
	}{
		{"testcase1", []int{1, 2, 4, 4, 4, 4}, 1, 3, true},
		{"testcase2", []int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2, true},
		{"testcase3", []int{1, 2, 1, 2, 1, 3}, 2, 3, false},
		{"testcase4", []int{1, 2, 3, 1, 2}, 2, 3, false},
		{"testcase5", []int{2, 2, 2, 2}, 2, 3, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := havePattern(tt.arr, tt.m, tt.k)
			if out != tt.exp {
				t.Fatalf("with input arr:%v m:%d, k:%d, wanted %t but got %t", tt.arr, tt.m, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
