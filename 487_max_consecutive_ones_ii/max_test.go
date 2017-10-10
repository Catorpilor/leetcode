package max

import "testing"

func TestMaxConsecutiveOnes(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single 1", []int{1}, 1},
		{"single 0", []int{0}, 1},
		{"all zeros", []int{0, 0, 0}, 1},
		{"10110", []int{1, 0, 1, 1, 0}, 4},
		{"1011011", []int{1, 0, 1, 1, 0, 1, 1}, 5},
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

func TestMaxConsecutiveOnes2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single 1", []int{1}, 1},
		{"single 0", []int{0}, 1},
		{"all zeros", []int{0, 0, 0}, 1},
		{"10110", []int{1, 0, 1, 1, 0}, 4},
		{"1011011", []int{1, 0, 1, 1, 0, 1, 1}, 5},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxConsecutiveOnes2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
