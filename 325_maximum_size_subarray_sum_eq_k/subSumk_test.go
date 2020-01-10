package subSumk

import "testing"

func TestMaxSubArrayLen(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 5, 0},
		{"k to large", []int{1, 2, 3}, 8, 0},
		{"testcase1", []int{1, -1, 5, -2, 3}, 3, 4},
		{"testcaes2", []int{-2, -1, 2, 1}, 1, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxSubLen(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k: %d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestBruteForce(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 5, 0},
		{"k to large", []int{1, 2, 3}, 8, 0},
		{"testcase1", []int{1, -1, 5, -2, 3}, 3, 4},
		{"testcaes2", []int{-2, -1, 2, 1}, 1, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := bruteForce(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k: %d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
