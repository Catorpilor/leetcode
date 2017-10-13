package long

import "testing"

func TestLengthOfLIS2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty nums", []int{}, 0},
		{"single element", []int{1}, 1},
		{"identical elements", []int{1, 1, 1, 1, 1}, 1},
		{"descreasing order", []int{4, 3, 2, 1}, 1},
		{"increasing order", []int{1, 2, 3, 4}, 4},
		{"normal case", []int{10, 9, 2, 5, 3, 7, 101, 10}, 4},
		{"test 1", []int{4, 10, 4, 3, 8, 9}, 3},
		{"test 2", []int{1, 3, 6, 7, 9, 4, 10, 5, 6}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LengthOfLIS2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
