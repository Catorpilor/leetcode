package ms

import "testing"

func TestMissingNumber(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"missing 2", []int{0, 1, 3}, 2},
		{"missing 3", []int{0, 1, 2, 4}, 3},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MissingNumber(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
