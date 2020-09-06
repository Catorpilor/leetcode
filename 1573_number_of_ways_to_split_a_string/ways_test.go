package ways

import "testing"

func TestNumOfWays(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty", "", 0},
		{"len(s)<4", "011", 0},
		{"testcase1", "0000", 3},
		{"testcase2", "10101", 4},
		{"testcase3", "100100010100110", 12},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfWays(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %t but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
