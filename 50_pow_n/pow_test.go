package pow

import (
	"math"
	"testing"
)

func TestPowN(t *testing.T) {
	st := []struct {
		name string
		x    float64
		n    int
		exp  float64
	}{
		{"x=0", 0.0, 123, 0.0},
		{"x=1", 1, 1234, 1},
		{"n is 2^-31", 2.0, math.MinInt32, 0.0},
		{"testcase1", 2.0, 10, 1024.0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := powN(tt.x, tt.n)
			if out != tt.exp {
				t.Fatalf("with input x:%f, x:%d wanted %02f, but got %02f", tt.x, tt.n, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
