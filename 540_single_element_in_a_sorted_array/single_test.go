package single

import "testing"

func TestSingleDup(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 1},
		{"testcase1", []int{1, 1, 2}, 2},
		{"testcase2", []int{1, 2, 2}, 1},
		{"testcase3", []int{1, 1, 2, 2, 3}, 3},
		{"testcase4", []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7}, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := singleDup(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
