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
		{"tle1", -1000000000, 44723},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := reachNumber(tt.target)
			if out != tt.exp {
				t.Fatalf("with input target: %d wanted %d but got %d", tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
