package uncrossed

import "testing"

func TestMaxLines(t *testing.T) {
	st := []struct {
		name       string
		arr1, arr2 []int
		exp        int
	}{
		{"both empty", nil, nil, 0},
		{"testcase1", []int{1, 4, 2}, []int{1, 2, 4}, 2},
		{"testcase2", []int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}, 3},
		{"testcase3", []int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxLines(tt.arr1, tt.arr2)
			if out != tt.exp {
				t.Fatalf("with input arr1:%v and arr2:%v wanted %d but got %d", tt.arr1, tt.arr2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
