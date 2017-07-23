package length

import "testing"

func TestLond(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		expect int
	}{
		{"nil slice", nil, 0},
		{"single element", []int{1}, 1},
		{"decreasing two elements", []int{2, 1}, 1},
		{"increasing two elements", []int{1, 2}, 2},
		{"ragular slice", []int{5, 3, 4, 8, 6, 7}, 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Lond(c.input)
			if ret != c.expect {
				t.Fatalf("expected %d but got %d ,with input %v",
					c.expect, ret, c.input)
			}
		})
	}
	t.Log("pass")
}
