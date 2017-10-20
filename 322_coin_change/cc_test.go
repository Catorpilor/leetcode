package cc

import "testing"

func TestCoinChange(t *testing.T) {
	st := []struct {
		name        string
		coins       []int
		amount, exp int
	}{
		{"no coins", nil, 1, -1},
		{"1 coin lower than amount", []int{2}, 3, -1},
		{"1 coin larger than amount", []int{3}, 2, -1},
		{"coins 1,2,5 with amount 11", []int{1, 2, 5}, 11, 3},
		{"coins 1,2,5 with amount 10", []int{1, 2, 5}, 10, 2},
		{"failed1", []int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CoinChange(c.coins, c.amount)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.coins, c.amount)
			}
		})
	}
}

func TestCoinChange2(t *testing.T) {
	st := []struct {
		name        string
		coins       []int
		amount, exp int
	}{
		{"no coins", nil, 1, -1},
		{"1 coin lower than amount", []int{2}, 3, -1},
		{"1 coin larger than amount", []int{3}, 2, -1},
		{"coins 1,2,5 with amount 11", []int{1, 2, 5}, 11, 3},
		{"coins 1,2,5 with amount 10", []int{1, 2, 5}, 10, 2},
		{"failed1", []int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CoinChange2(c.coins, c.amount)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v and %d",
					c.exp, ret, c.coins, c.amount)
			}
		})
	}
}
