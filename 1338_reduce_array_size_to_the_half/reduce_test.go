package reduce

import "testing"

func TestMinSetSize(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"single one", []int{1}, 1},
		{"all identical", []int{7, 7, 7, 7, 7}, 1},
		{"testcase1", []int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{"testcase2", []int{1, 9}, 1},
		{"testcase3", []int{1000, 1000, 3, 7}, 1},
		{"testcase4", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minSetSize(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr: %v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
