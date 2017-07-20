package robber

import "testing"

func TestRob(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{"nil slice", nil, 0},
		{"single element", []int{1}, 1},
		{"2 elements", []int{1, 2}, 2},
		{"3 elements", []int{2, 1, 4}, 6},
		{"5 elements", []int{2, 1, 4, 3, 7}, 13},
		{"6 elements", []int{2, 1, 4, 3, 7, 5}, 13},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Rob(c.nums)
			if c.expect != ret {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.expect, ret, c.nums)
			}
		})
	}
	t.Log("pass")
}

func TestRobSimplify(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{"nil slice", nil, 0},
		{"single element", []int{1}, 1},
		{"2 elements", []int{1, 2}, 2},
		{"3 elements", []int{2, 1, 4}, 6},
		{"5 elements", []int{2, 1, 4, 3, 7}, 13},
		{"6 elements", []int{2, 1, 4, 3, 7, 5}, 13},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Rob(c.nums)
			if c.expect != ret {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.expect, ret, c.nums)
			}
		})
	}
	t.Log("pass")
}
