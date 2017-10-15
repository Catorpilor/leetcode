package average

import "testing"

func TestMaxAverage(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  float64
	}{
		{"empty slice", []int{}, 4, 0},
		{"max value", []int{1, 2, 3, 4, 5}, 1, 5.0},
		{"example", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxAverage(c.nums, c.k)
			if ret != c.exp {
				t.Fatalf("expected %.2f but got %.2f with inputs %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
