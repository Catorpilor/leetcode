package sq

import "testing"

func TestLengthOfLCIS(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty nums", []int{}, 0},
		{"single element", []int{1}, 1},
		{"identical elements", []int{2, 2, 2, 2}, 1},
		{"13547", []int{1, 3, 5, 4, 7}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LengthOfLCIS(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
