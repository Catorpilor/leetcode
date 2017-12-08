package pos

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 1},
		{"all negative", []int{-1, -2}, 1},
		{"all identical", []int{2, 2, 2, 2}, 1},
		{"test 1", []int{1, 2, 0}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FirstMissingPositive(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}

func TestFirstMissingPositive2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 1},
		{"all negative", []int{-1, -2}, 1},
		{"all identical", []int{2, 2, 2, 2}, 1},
		{"test 1", []int{1, 2, 0}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FirstMissingPositive2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.nums)
			}
		})
	}
}
