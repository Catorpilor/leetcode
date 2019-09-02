package subarr

import "testing"

func TestSubarraysWithKDistinct(t *testing.T) {
	st := []struct {
		name  string
		input []int
		k     int
		exp   int
	}{
		{"nil slice", nil, 1, 0},
		{"single element", []int{1}, 1, 1},
		{"k is larger than the slice", []int{1, 2, 2}, 4, 0},
		{"distinct num is less than k", []int{1, 2, 1, 2, 3}, 5, 0},
		{"testcase1", []int{1, 2, 1, 2, 3}, 2, 7},
		{"testcase2", []int{1, 2, 1, 3, 4}, 3, 3},
		{"testcase3", []int{1, 2, 1, 3}, 2, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := subarraysWithKDistinct(tt.input, tt.k)
			if out != tt.exp {
				t.Fatalf("with input: %v and k: %d wanted %d but got %d\n", tt.input, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
