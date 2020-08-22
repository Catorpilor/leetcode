package choc

import "testing"

func TestDivideChocolate(t *testing.T) {
	st := []struct {
		name      string
		chocolate []int
		k         int
		exp       int
	}{
		{"testcase1", []int{6, 3, 2, 8, 7, 5}, 3, 9},
		{"testcase2", []int{6, 3, 2, 8, 7, 5}, 6, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := divideChocolate(tt.chocolate, tt.k)
			if out != tt.exp {
				t.Fatalf("with input chocolate:%v and k:%d wanted %d but got %d", tt.chocolate, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
