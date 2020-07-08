package mvoe

import "testing"

func TestFindMaxValueOfEquation(t *testing.T) {
	st := []struct {
		name   string
		points [][]int
		k, exp int
	}{
		{"testcase1", [][]int{[]int{1, 3}, []int{2, 0}, []int{5, 10}, []int{6, -10}}, 1, 4},
		{"testcase2", [][]int{[]int{0, 0}, []int{3, 0}, []int{9, 2}}, 3, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findMaxValueOfEquation(tt.points, tt.k)
			if out != tt.exp {
				t.Fatalf("with input points:%v and k:%d wanted %d but got %d", tt.points, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestFindMaxValueOfEquationUseDeque(t *testing.T) {
	st := []struct {
		name   string
		points [][]int
		k, exp int
	}{
		{"testcase1", [][]int{[]int{1, 3}, []int{2, 0}, []int{5, 10}, []int{6, -10}}, 1, 4},
		{"testcase2", [][]int{[]int{0, 0}, []int{3, 0}, []int{9, 2}}, 3, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useDeque(tt.points, tt.k)
			if out != tt.exp {
				t.Fatalf("with input points:%v and k:%d wanted %d but got %d", tt.points, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
