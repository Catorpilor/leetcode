package splk

import "testing"

func TestNumsOfSubarray(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 100, 0},
		{"not qulified", []int{3, 6, 9}, 2, 0},
		{"testcase1", []int{10, 5, 2, 6}, 100, 8},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numsOfSubarray(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with nums: %v and k: %d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
