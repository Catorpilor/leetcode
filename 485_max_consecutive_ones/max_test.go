package max

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty nums", []int{}, 0},
		{"no 1", []int{0, 0, 0}, 0},
		{"max is 3", []int{1, 1, 0, 1, 1, 1}, 3},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxConsecutiveOnes(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
