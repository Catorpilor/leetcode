package vtn

import "testing"

func TestTriangleNumbers(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"not valid", []int{0, 0, 0}, 0},
		{"testcase1", []int{2, 2, 3, 4}, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := triangleNumbers(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestUseBs(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"not valid", []int{0, 0, 0}, 0},
		{"testcase1", []int{2, 2, 3, 4}, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useBs(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
