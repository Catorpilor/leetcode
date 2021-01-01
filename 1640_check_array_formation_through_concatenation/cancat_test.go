package concat

import "testing"

func TestCanFormArray(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		pcs  [][]int
		exp  bool
	}{
		{"testcase1", []int{85}, [][]int{[]int{85}}, true},
		{"testcase2", []int{15, 88}, [][]int{[]int{88}, []int{15}}, true},
		{"testcase3", []int{49, 18, 16}, [][]int{[]int{16, 18, 49}}, false},
		{"testcase4", []int{91, 4, 64, 78}, [][]int{[]int{78}, []int{4, 64}, []int{91}}, true},
		{"testcase5", []int{1, 3, 5, 7}, [][]int{[]int{2, 4, 6, 8}}, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := canFormArray(tt.arr, tt.pcs)
			if out != tt.exp {
				t.Fatalf("with input arr:%v pcs:%v wanted %t but got %t", tt.arr, tt.pcs, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
