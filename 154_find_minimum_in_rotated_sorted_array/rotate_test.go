package rotate

import "testing"

func TestFindMinimum(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single element", []int{1}, 1},
		{"all identical", []int{1, 1, 1, 1, 1, 1}, 1},
		{"rotated at 4", []int{4, 5, 6, 7, 0, 1, 2, 3, 4}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindMinimum(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestFindMinimum2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"single element", []int{1}, 1},
		{"all identical", []int{1, 1, 1, 1, 1, 1}, 1},
		{"rotated at 4", []int{4, 5, 6, 7, 0, 1, 2, 3, 4}, 0},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindMinimum2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
