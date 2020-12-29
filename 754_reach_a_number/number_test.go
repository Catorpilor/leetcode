package number

import "testing"

func TestReachNumber(t *testing.T) {
	st := []struct {
		name   string
		target int
		exp    int
	}{
		{"testcase1", 3, 2},
		{"testcase2", 2, 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reachNumber(tt.target)
			if out != tt.exp {
				t.Fataf("with input target: %d wanted %d", tt.target, tt.exp)
			}
			t.Log("pass")
		})
	}
}
