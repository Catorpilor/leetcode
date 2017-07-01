package searchinsertposition

import "testing"

type res struct {
	nums   []int
	target int
	expect int
}

func TestSearchInsert(t *testing.T) {
	cases := []*res{
		&res{nil, 1, 0},
		&res{[]int{1, 2, 3}, 4, 3},
		&res{[]int{1, 3, 5, 6}, 5, 2},
		&res{[]int{1, 3, 5, 6}, 7, 4},
	}
	for _, c := range cases {
		ret := SearchInsert(c.nums, c.target)
		if ret != c.expect {
			t.Fatalf("expected %d, but got %d, with inputs %v and target %d",
				c.expect, ret, c.nums, c.target)
		}
	}

	t.Log("pass")
}
