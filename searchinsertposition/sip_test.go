package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		expect int
	}{
		{"nil nums", nil, 1, 0},
		{"normal 123", []int{1, 2, 3}, 4, 3},
		{"insert duplicate", []int{1, 3, 5, 6}, 5, 2},
		{"insert normal", []int{1, 3, 5, 6}, 7, 4},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := SearchInsert(c.nums, c.target)
			if ret != c.expect {
				t.Fatalf("expected %d, but got %d, with inputs %v and target %d",
					c.expect, ret, c.nums, c.target)
			}
		})
	}

	t.Log("pass")
}
