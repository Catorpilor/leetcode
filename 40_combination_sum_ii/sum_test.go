package sum

import "testing"

func TestCombinationSum(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		target int
		exp    [][]int
	}{
		{"single", []int{1}, 6, nil},
		{"two", []int{1, 2}, 7, nil},
		{"test1", []int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{
			[]int{1, 7}, []int{1, 2, 5}, []int{2, 6}, []int{1, 1, 6},
		}},
		{"failed 1", []int{2, 2, 2}, 4, [][]int{[]int{2, 2}}},
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
