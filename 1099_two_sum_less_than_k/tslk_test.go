package tslk

import "testing"

func TestTwoSumLessThanK(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 3, -1},
		{"less than 2", []int{1}, 2, -1},
		{"bigger than k", []int{2, 3}, 4, -1},
		{"testcase1", []int{34, 23, 1, 24, 75, 33, 54, 8}, 60, 58},
		{"testcase2", []int{10, 20, 30}, 15, -1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := twoSumLessThanK(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
