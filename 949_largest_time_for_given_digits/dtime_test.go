package dtime

import "testing"

func TestMaxTime(t *testing.T) {
	st := []struct {
		name string
		arr  []int
		exp  string
	}{
		{"all 0s", []int{0, 0, 0, 0}, "00:00"},
		{"testcase1", []int{1, 2, 3, 4}, "23:41"},
		{"testcase2", []int{3, 3, 3, 3}, ""},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxTime(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %s but got %s", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
