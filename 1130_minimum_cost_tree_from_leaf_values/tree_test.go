package tree

import "testing"

func TestMCTFromLeafValues(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"testcase1", []int{6, 2, 4}, 32},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := mctFromLeafValues(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("Pass")
		})
	}
}

func TestMCTFromLeafValuesUseGreedy(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"testcase1", []int{6, 2, 4}, 32},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useGreedy(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("Pass")
		})
	}
}

func TestMCTFromLeafValuesUseStack(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  int
	}{
		{"testcase1", []int{6, 2, 4}, 32},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useStack(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("Pass")
		})
	}
}
