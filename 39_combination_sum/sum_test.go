package sum

import "testing"

func TestCombinationSum(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    [][]int
	}{
		{"single", []int{1}, 6, [][]int{[]int{1, 1, 1, 1, 1, 1}}},
		{"two", []int{1, 2}, 7, [][]int{[]int{1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 2, 2}, []int{1, 2, 2, 2},
			[]int{1, 1, 1, 1, 1, 2}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CombinationSum(c.nums, c.target)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v, with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}
}
