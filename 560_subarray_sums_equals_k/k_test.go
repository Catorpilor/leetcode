package sum

import "testing"

func TestSubArraySum(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 1, 0},
		{"single element", []int{1}, 1, 1},
		{"identical", []int{1, 1, 1}, 3, 1},
		{"failed 1", []int{28, 54, 7, -70, 22, 65, -6}, 100, 1},
		{"failed 2", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 55},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SubArraySum(c.nums, c.k)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}

func TestSubArraySum2(t *testing.T) {
	st := []struct {
		name   string
		nums   []int
		k, exp int
	}{
		{"empty slice", []int{}, 1, 0},
		{"single element", []int{1}, 1, 1},
		{"identical", []int{1, 1, 1}, 3, 1},
		{"failed 1", []int{28, 54, 7, -70, 22, 65, -6}, 100, 1},
		{"failed 2", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 55},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SubArraySum2(c.nums, c.k)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
