package di

import "testing"

func TestDivide(t *testing.T) {
	st := []struct {
		name              string
		dividend, divisor int
		exp               int
	}{
		{"testcase1", 10, 3, 3},
		{"testcase2", 7, -3, -2},
		{"testcase3", 0, 12, 0},
		{"testcase4", 1, 1, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := divide(tt.dividend, tt.divisor)
			if out != tt.exp {
				t.Fatalf("wanted %d  but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
