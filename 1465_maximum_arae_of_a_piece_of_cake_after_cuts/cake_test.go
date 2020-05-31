package cake

import "testing"

func TestMaxAera(t *testing.T) {
	st := []struct {
		name       string
		w, h       int
		horizontal []int
		vertical   []int
		exp        int
	}{
		{"testcase1", 5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
		{"testcase2", 5, 4, []int{3, 1}, []int{1}, 9},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxArea(tt.w, tt.h, tt.horizontal, tt.vertical)
			if out != tt.exp {
				t.Fatalf("with input w:%d, h:%d, hori:%v, vertical:%v wanted %d but got %d", tt.w, tt.h, tt.horizontal, tt.vertical, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
