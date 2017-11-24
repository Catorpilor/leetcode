package sum

import "testing"

func TestThreeSumClosest(t *testing.T) {
	st := []struct {
		name        string
		nums        []int
		target, exp int
	}{
		{"3 elements", []int{1, 2, 3}, 1, 6},
		{"4 elements", []int{1, -1, 2, -4}, 1, 2},
		{"identical elementes", []int{1, 1, 1, 1}, 1, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ThreeSumClosest(c.nums, c.target)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.nums, c.target)
			}
		})
	}

}
