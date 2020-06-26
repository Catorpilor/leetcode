package single

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"testcase1", []int{0, 0, 1, 0}, 1},
		{"testcase2", []int{0, 1, 2, 0, 2, 2, 1, 1, 0, 3}, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := singleNumber(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %d but got %d", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
