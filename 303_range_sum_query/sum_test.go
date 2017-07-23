package sum

import "testing"

var nums = []int{-2, 0, 3, -5, 2, -1}

var na = NewNumArray(nums)

func TestSumRange(t *testing.T) {
	cases := []struct {
		name   string
		i      int
		j      int
		expect int
	}{
		{"i=j=0", 0, 0, -2},
		{"i=0,j=2", 0, 2, 1},
		{"2,5", 2, 5, -1},
		{"0,5", 0, 5, -3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := na.SumRange(c.i, c.j)
			if ret != c.expect {
				t.Fatalf("expect %d, but got %d, with input %d, %d",
					c.expect, ret, c.i, c.j)
			}
		})
	}
	t.Log("pass")
}
