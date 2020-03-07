package single

import "testing"

func TestSingleNumber(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"testcase1", []int{1, 1, 2, 3, 2}, 3},
		{"testcase2", []int{2, 3, 4, 2, 3}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := singleNum(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr: %v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
