package rect

import "testing"

func initWithTesting(rects [][]int, t *testing.T) *Solution {
	s := Constructor(rects)
	t.Cleanup(func() {
		s = nil
	})
	return s
}

func TestPick(t *testing.T) {
	s := initWithTesting([][]int{[]int{1, 1, 5, 5}}, t)
	cmp := func(p []int) bool {
		return p[0] >= 1 && p[0] <= 5 && p[1] >= 1 && p[1] <= 5
	}
	st := []struct {
		name string
		exp  func([]int) bool
	}{
		{"pick 1", cmp},
		{"pick 2", cmp},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := s.Pick()
			if !cmp(out) {
				t.Fatalf("should in the bound but got %v", out)
			}
			t.Log("pass")
		})
	}
}
