package stack

import "testing"

func TestIsValid(t *testing.T) {
	st := []struct {
		name      string
		push, pop []int
		exp       bool
	}{
		{"testcase1", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{"testcase2", []int{1, 2, 3, 4, 5}, []int{3, 5, 4, 1, 2}, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isValid(tt.push, tt.pop)
			if out != tt.exp {
				t.Fatalf("wanted %t but got %t", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
