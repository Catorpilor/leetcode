package max

import "testing"

func TestMaxArea(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"test 1", []int{2, 1, 3, 5, 4}, 8},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxArea(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}

func TestMaxArea2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"test 1", []int{2, 1, 3, 5, 4}, 8},
		{"failed 1", []int{2, 3, 4, 5, 18, 17, 6}, 17},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxArea2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
