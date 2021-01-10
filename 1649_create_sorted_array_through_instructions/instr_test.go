package inster

import "testing"

func TestCreateSortedArray(t *testing.T) {
	st := []struct {
		name         string
		instructions []int
		exp          int
	}{
		{"testcase1", []int{1, 5, 6, 2}, 1},
		{"testcase2", []int{1, 2, 3, 6, 5, 4}, 3},
		{"testcase3", []int{1, 3, 3, 3, 2, 4, 2, 1, 2}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := createSortedArray(tt.instructions)
			if out != tt.exp {
				t.Fatalf("with input instructions:%v wanted %d but got %d", tt.instructions, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
