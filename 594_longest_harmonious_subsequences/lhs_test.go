package lhs

import "testing"

func TestFindLHS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"single element", []int{1}, 0},
		{"all identical", []int{1, 1, 1, 1}, 0},
		{"testcase1", []int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findLHS(tt.nums)
			if out != tt.exp {
				t.Fatalf("with nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
