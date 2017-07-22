package sell

import "testing"

func TestSell(t *testing.T) {
	cases := []struct {
		name   string
		prices []int
		expect int
	}{
		{"nil input", nil, 0},
		{"one element", []int{1}, 0},
		{"two elements", []int{1, 2}, 1},
		{"decreasing order", []int{3, 2, 1}, 0},
		{"increasing order", []int{1, 2, 3}, 2},
		{"random elements", []int{7, 1, 5, 3, 6, 4}, 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Sell(c.prices)
			if ret != c.expect {
				t.Fatalf("expected %d, but  got %d, with input %v",
					c.expect, ret, c.prices)
			}
		})
	}
	t.Log("pass")
}
