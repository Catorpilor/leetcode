package target

import "testing"

func TestMinLength(t *testing.T) {
	st := []struct {
		name   string
		arr    []int
		target int
		exp    int
	}{
		{"testcase1", []int{3, 2, 2, 4, 3}, 3, 2},
		{"testcase2", []int{7, 3, 4, 7}, 7, 2},
		{"testcase3", []int{4, 3, 2, 6, 2, 3, 4}, 6, -1},
		{"testcase4", []int{5, 5, 4, 4, 5}, 3, -1},
		{"testcase5", []int{3, 1, 1, 1, 5, 1, 2, 1}, 3, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minLength(tt.arr, tt.target)
			if out != tt.exp {
				t.Fatalf("with input arr:%v and target:%d wanted %d but got %d", tt.arr, tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
