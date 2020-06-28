package pairs

import "testing"

func TestCanArrange(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		k    int
		exp  bool
	}{
		{"odd ones", []int{1, 2, 3}, 2, false},
		{"k=1", []int{1, 2, 3, 4}, 1, true},
		{"testcase1", []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}, 5, true},
		{"testcase2", []int{1, 2, 3, 4, 5, 6}, 7, true},
		{"testcase3", []int{1, 2, 3, 4, 5, 6}, 10, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := canArrange(tt.arr, tt.k)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and k:%d wanted %t but got %t", tt.arr, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
