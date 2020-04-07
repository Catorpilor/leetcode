package count

import "testing"

func TestCounting(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"empty arr", nil, 0},
		{"single element", []int{1}, 0},
		{"all identical", []int{1, 1, 1}, 0},
		{"testcase1", []int{1, 2, 3}, 2},
		{"testcsae2", []int{1, 1, 1, 1, 1, 2}, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countElements(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
		})
	}
}
