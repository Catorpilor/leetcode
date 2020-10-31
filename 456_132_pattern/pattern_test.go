package pattern

import "testing"

func TestCanFindPattern(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"just one", []int{1}, false},
		{"testcase1", []int{1, 2, 3, 4, 5}, false},
		{"testcase2", []int{3, 1, 4, 2}, true},
		{"testcase3", []int{-1, 3, 2, 0}, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findPattern(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %t but got %t", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
