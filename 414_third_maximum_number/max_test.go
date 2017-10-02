package max

import "testing"

func TestThirdMax(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"len(nums) < 3", []int{1, 2}, 2},
		{"len(nums)=3 without thridmax", []int{3, 2, 2}, 3},
		{"normal case", []int{2, 2, 3, 1}, 1},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := ThirdMax(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with input %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
