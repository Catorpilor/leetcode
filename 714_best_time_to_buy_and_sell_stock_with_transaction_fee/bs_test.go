package bs

import "testing"

func TestMaxProfit(t *testing.T) {
	st := []struct {
		name     string
		prices   []int
		fee, exp int
	}{
		{"single element", []int{1}, 2, 0},
		{"two elements", []int{1, 2}, 2, 0},
		{"decreasing", []int{3, 2, 1}, 2, 0},
		{"increasing", []int{1, 2, 4}, 2, 1},
		{"test 1", []int{1, 3, 2, 8, 4, 9}, 2, 8},
		{"failed 1", []int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxProfit(c.prices, c.fee)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.prices, c.fee)
			}
		})
	}
}

func TestMaxProfit2(t *testing.T) {
	st := []struct {
		name     string
		prices   []int
		fee, exp int
	}{
		{"single element", []int{1}, 2, 0},
		{"two elements", []int{1, 2}, 2, 0},
		{"decreasing", []int{3, 2, 1}, 2, 0},
		{"increasing", []int{1, 2, 4}, 2, 1},
		{"test 1", []int{1, 3, 2, 8, 4, 9}, 2, 8},
		{"failed 1", []int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxProfit2(c.prices, c.fee)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.prices, c.fee)
			}
		})
	}
}
