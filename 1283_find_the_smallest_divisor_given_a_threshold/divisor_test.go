package divisor

import "testing"

func TestSmallestDivisor(t *testing.T) {
	st := []struct {
		name      string
		nums      []int
		threshold int
		exp       int
	}{
		{"testcase1", []int{19}, 5, 4},
		{"testcase2", []int{1, 2, 5, 9}, 6, 5},
		{"testcase3", []int{2, 3, 5, 7, 11}, 11, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := smallestDivisor(tt.nums, tt.threshold)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and threshold:%d wanted %d but got %d", tt.nums, tt.threshold, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
