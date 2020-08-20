package balls

import "testing"

func TestMaxDistances(t *testing.T) {
	st := []struct {
		name     string
		position []int
		m        int
		exp      int
	}{
		{"testcase1", []int{1, 2, 3, 4, 7}, 3, 3},
		{"testcase2", []int{5, 4, 3, 2, 1, 1000000000}, 2, 999999999},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxDistance(tt.position, tt.m)
			if out != tt.exp {
				t.Fatalf("with input position:%v and m:%d wanted %d but got %d", tt.position, tt.m, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
