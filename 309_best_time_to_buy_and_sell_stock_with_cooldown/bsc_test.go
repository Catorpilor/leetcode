package bsc

import "testing"

func TestMaxProfit(t *testing.T) {
	// t.Fatal("not implemented")
	st := []struct {
		name   string
		prices []int
		exp    int
	}{
		{"nil slice", nil, 0},
		{"single element", []int{1}, 0},
		{"two elements with descending order", []int{2, 1}, 0},
		{"two elements with ascending order", []int{1, 3}, 2},
		{"test case", []int{1, 2, 3, 0, 2}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxProfit(c.prices)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.prices)
			}
		})
	}
}
