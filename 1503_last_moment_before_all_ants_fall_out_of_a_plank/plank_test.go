package plank

import "testing"

func TestLastMoment(t *testing.T) {
	st := []struct {
		name        string
		n           int
		left, right []int
		exp         int
	}{
		{"left is nil", 5, nil, []int{3}, 2},
		{"right is nil", 5, []int{3}, nil, 3},
		{"testcase1", 4, []int{4, 3}, []int{0, 1}, 4},
		{"testcase2", 9, []int{5}, []int{4}, 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := lastMoment(tt.n, tt.left, tt.right)
			if out != tt.exp {
				t.Fatalf("with input n:%d, left:%v and right:%v wanted %d but got %d", tt.n, tt.left, tt.right, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
