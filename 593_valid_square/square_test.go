package square

import "testing"

func TestIsValid(t *testing.T) {
	st := []struct {
		name           string
		p1, p2, p3, p4 []int
		exp            bool
	}{
		{"testcase1", []int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, true},
		{"testcase2", []int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}, true},
		{"testcase3", []int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isValid(tt.p1, tt.p2, tt.p3, tt.p4)
			if out != tt.exp {
				t.Fatalf("wanted %t but got %t", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
