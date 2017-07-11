package major

import (
	"testing"
)

func TestMajority(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		expect int
	}{
		{"many 1s", []int{1, 2, 1}, 1},
		{"1 element", []int{1}, 1},
		{"2 elements", []int{1, 2}, 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := Majority(c.nums)
			if ret != c.expect {
				t.Fatalf("expected %d, but got %d, with nums %v",
					c.expect, ret, c.nums)
			}
		})
	}
	t.Log("pass")
}
