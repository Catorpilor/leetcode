package paint

import "testing"

func TestMinCost(t *testing.T) {
	cases := []struct {
		name   string
		costs  [][]int
		expect int
	}{
		{"nil slice", nil, 0},
		{"1 row", [][]int{[]int{1, 2, 3}}, 1},
		{"normal case", [][]int{[]int{2, 1, 4}, []int{4, 1, 2}, []int{2, 5, 1}, []int{3, 4, 2}}, 7},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := MinCost(c.costs)
			if ret != c.expect {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.expect, ret, c.costs)
			}
		})
	}
	t.Log("pass")
}
