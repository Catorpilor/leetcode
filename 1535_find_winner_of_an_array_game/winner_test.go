package winner

import "testing"

func TestWinnerOfGame(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		k    int
		exp  int
	}{
		{"k=1", []int{3, 1, 2, 4, 5}, 1, 3},
		{"testcase1", []int{2, 1, 3, 5, 4, 6, 7}, 2, 5},
		{"testcase2", []int{3, 2, 1}, 10, 3},
		{"testcase3", []int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7, 9},
		{"testcase4", []int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 100000000, 99},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useOnePassWithStack(tt.arr, tt.k)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and k:%d wanted %d but got %d", tt.arr, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
