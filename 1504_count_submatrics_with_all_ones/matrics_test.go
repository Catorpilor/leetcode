package matrics

import "testing"

func TestCountMatrics(t *testing.T) {
	st := []struct {
		name string
		mat  [][]int
		exp  int
	}{
		{"empty mat", nil, 0},
		{"testcase1", [][]int{[]int{1, 1, 1, 1, 1, 1}}, 21},
		{"testcase2", [][]int{[]int{1, 0, 1}, []int{1, 1, 0}, []int{1, 1, 0}}, 13},
		{"testcase3", [][]int{[]int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 1, 1, 0}}, 24},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countMatrics(tt.mat)
			if out != tt.exp {
				t.Fatalf("with input mat:%v, wanted %d, but got %d", tt.mat, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestCountMatricsUseStack(t *testing.T) {
	st := []struct {
		name string
		mat  [][]int
		exp  int
	}{
		{"empty mat", nil, 0},
		{"testcase1", [][]int{[]int{1, 1, 1, 1, 1, 1}}, 21},
		{"testcase2", [][]int{[]int{1, 0, 1}, []int{1, 1, 0}, []int{1, 1, 0}}, 13},
		{"testcase3", [][]int{[]int{0, 1, 1, 0}, []int{0, 1, 1, 1}, []int{1, 1, 1, 0}}, 24},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useStack(tt.mat)
			if out != tt.exp {
				t.Fatalf("with input mat:%v, wanted %d, but got %d", tt.mat, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
