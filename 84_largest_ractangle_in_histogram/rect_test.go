package rect

import "testing"

func TestLargestRectangle(t *testing.T) {
	st := []struct {
		name    string
		heights []int
		exp     int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"identical", []int{5, 5}, 10},
		{"test 1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"test 2", []int{2, 1, 2, 3, 1}, 5},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LargestRectangle(c.heights)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.heights)
			}
		})
	}
}

func TestLargestRectangle2(t *testing.T) {
	st := []struct {
		name    string
		heights []int
		exp     int
	}{
		{"empty", []int{}, 0},
		{"single", []int{1}, 1},
		{"identical", []int{5, 5}, 10},
		{"test 1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"test 2", []int{2, 1, 2, 3, 1}, 5},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LargestRectangle2(c.heights)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v", c.exp, ret, c.heights)
			}
		})
	}
}
