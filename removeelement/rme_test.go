package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		name   string
		inputs []int
		val    int
		expect int
	}{
		{"normal", []int{3, 2, 2, 3}, 3, 2},
		{"nil inputs", nil, 3, 0},
		{"inputs are same", []int{3, 3, 3}, 3, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := RemoveElement(c.inputs, c.val)
			if ret != c.expect {
				t.Fatalf("expected %d, but got %d, with inputs: %v, and val: %d",
					c.expect, ret, c.inputs, c.val)
			}
		})
	}

	t.Log("passed")
}
