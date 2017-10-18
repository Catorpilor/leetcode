package msa

import "testing"

func TestMaxSubArray(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{2}, 2},
		{"two elmements", []int{1, -2}, 1},
		{"both negative", []int{-1, -1, -1}, -1},
		{"test case", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxSubArray(c.nums)
			if ret != c.exp {
				t.Fatalf("expectde %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
