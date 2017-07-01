package removeelement

import "testing"

type res struct {
	inputs []int
	val    int
	expect int
}

func TestRemoveElement(t *testing.T) {
	cases := []*res{
		&res{[]int{3, 2, 2, 3}, 3, 2},
		&res{nil, 3, 0},
		&res{[]int{3, 3, 3}, 3, 0},
	}

	for _, c := range cases {
		ret := RemoveElement(c.inputs, c.val)
		if ret != c.expect {
			t.Fatalf("expected %d, but got %d, with inputs: %v, and val: %d",
				c.expect, ret, c.inputs, c.val)
		}
	}

	t.Log("passed")
}
