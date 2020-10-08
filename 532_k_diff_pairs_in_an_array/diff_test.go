package diff

import "testing"

func TestFindPairs(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  int
	}{
		{"testcase1", []int{3, 1, 4, 1, 5}, 2, 2},
		{"testcase2", []int{1, 2, 3, 4, 5}, 1, 4},
		{"testcase3", []int{1, 3, 1, 4, 5}, 0, 1},
		{"testcase4", []int{1, 2, 4, 4, 3, 3, 0, 9, 2, 3}, 3, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findPairs(tt.nums, tt.k)
			if out != tt.exp {
				t.Fatalf("with input nums:%v and k:%d wanted %d but got %d", tt.nums, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
