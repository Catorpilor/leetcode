package dup

import "testing"

func TestContainsDups(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		k, t int
		exp  bool
	}{
		{"testcase1", []int{1, 2, 3, 1}, 3, 0, true},
		{"testcase2", []int{1, 0, 1, 1}, 1, 2, true},
		{"testcase3", []int{1, 5, 9, 1, 5, 9}, 2, 3, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := containsDups(tt.arr, tt.k, tt.t)
			if out != tt.exp {
				t.Fatalf("with input arr:%v k:%d t:%d wanted %t but got %t", tt.arr, tt.k, tt.t, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
