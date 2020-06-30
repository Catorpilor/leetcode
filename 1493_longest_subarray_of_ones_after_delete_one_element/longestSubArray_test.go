package longestSubArray

import "testing"

func TestLongestSubArrayOfOnes(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"all ones", []int{1, 1, 1, 1}, 3},
		{"all zeros", []int{0, 0, 0, 0}, 0},
		{"single one", []int{1}, 0},
		{"testcaes1", []int{1, 1, 0, 1}, 3},
		{"testcase2", []int{0, 1, 1, 1, 0, 1, 1}, 5},
		{"testcase3", []int{1, 1, 0, 0, 1, 1, 1, 0, 1}, 4},
		{"failed1", []int{1, 0, 0, 1, 0}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestSub(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
