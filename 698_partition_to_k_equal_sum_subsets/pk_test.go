package pk

import "testing"

func TestCanPK(t *testing.T) {
	st := []struct {
		name string
		nums []int
		k    int
		exp  bool
	}{
		{"single slice with k > 1", []int{1}, 2, false},
		{"single slice with k=1", []int{2}, 1, true},
		{"test 1", []int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{"failed 1", []int{2, 2, 10, 5, 2, 7, 2, 2, 13}, 3, true},
		{"failed 2", []int{2, 2, 2, 2, 3, 4, 5}, 4, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CanPK(c.nums, c.k)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v and %d",
					c.exp, ret, c.nums, c.k)
			}
		})
	}
}
