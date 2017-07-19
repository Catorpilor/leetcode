package sum

import "testing"

func TestNumberOfCoins(t *testing.T) {
	cases := []struct {
		name   string
		coins  []int
		sum    int
		expect int
	}{
		{"# of 0", []int{1, 3, 5}, 0, 0},
		{"# of 1", []int{1, 3, 5}, 1, 1},
		{"# of 2", []int{1, 3, 5}, 2, 2},
		{"# of 3", []int{1, 3, 5}, 3, 1},
		{"# of 4", []int{1, 3, 5}, 4, 2},
		{"# of 5", []int{1, 3, 5}, 5, 1},
		{"# of 6", []int{1, 3, 5}, 6, 2},
		{"# of 7", []int{1, 3, 5}, 7, 3},
		{"# of 8", []int{1, 3, 5}, 8, 2},
		{"# of 9", []int{1, 3, 5}, 9, 3},
		{"# of 10", []int{1, 3, 5}, 10, 2},
		{"# of 11", []int{1, 3, 5}, 11, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := NumberOfCoins(c.coins, c.sum)
			if ret != c.expect {
				t.Fatalf("expected %d but got %d, with coins %v and sum %d",
					c.expect, ret, c.coins, c.sum)
			}
		})
	}
	t.Log("pass")
}
