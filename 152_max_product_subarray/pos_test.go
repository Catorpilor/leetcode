package pos

import "testing"

func TestMaxPos(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"single element", []int{1}, 1},
		{"two positive elements", []int{2, 3}, 6},
		{"two with one negtive", []int{2, -3}, 2},
		{"failed1", []int{7, -2, -4}, 56},
		{"failed2", []int{2, -5, -2, -4, 3}, 24},
		{"test1", []int{2, 3, -2, 4}, 6},
		{"test2", []int{-2, -3, -3, 2, 4, -4}, 576},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxPos(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
