package pot

import "testing"

func TestMaxPot(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"length of 3", []int{3, 2, 1}, 6},
		{"all positive", []int{5, 4, 3, 2, 1}, 60},
		{"with negtive", []int{-5, -4, 1, 2, 3}, 60},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxPot(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestMaxPot2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"length of 3", []int{3, 2, 1}, 6},
		{"all positive", []int{5, 4, 3, 2, 1}, 60},
		{"with negtive", []int{-5, -4, 1, 2, 3}, 60},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxPot2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
