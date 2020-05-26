package dot

import "testing"

func TestMaxDotProduct(t *testing.T) {
	st := []struct {
		name       string
		num1, num2 []int
		exp        int
	}{
		{"all single", []int{-1}, []int{1}, -1},
		{"testcase1", []int{-1, -1}, []int{1, 1}, -1},
		{"testcase2", []int{3, -2}, []int{2, -6, 7}, 21},
		{"testcase3", []int{2, 1, -2, 5}, []int{3, 0, 6}, 18},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDotProduct(tt.num1, tt.num2)
			if out != tt.exp {
				t.Fatalf("with input num1:%v and num2:%v wanted %d but got %d", tt.num1, tt.num2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
