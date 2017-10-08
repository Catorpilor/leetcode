package move

import "testing"

func TestMinMoves2(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{1}, 0},
		{"123", []int{1, 2, 3}, 3},
		{"56885", []int{5, 6, 8, 8, 5}, 7},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MinMoves2(c.nums)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d, with inputs %v",
					c.exp, ret, c.nums)
			}
		})
	}
}
