package sum

import "testing"

func TestCombinationSum(t *testing.T) {
	st := []struct {
		name string
		n, k int
		exp  [][]int
	}{
		{"n is 0", 0, 2, nil},
		{"k is 0", 2, 0, nil},
		{"k3n7", 7, 3, [][]int{[]int{1, 2, 4}}},
		{"k3n9", 9, 3, [][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := combinationSum(c.k, c.n)
			if len(ret) != len(c.exp) {
				t.Fatalf("expected %v but got %v, with input %d and %d",
					c.exp, ret, c.k, c.n)
			}
		})
	}
}
