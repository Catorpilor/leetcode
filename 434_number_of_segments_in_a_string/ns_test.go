package ns

import "testing"

func TestCountSegments(t *testing.T) {
	st := []struct {
		name, s string
		exp     int
	}{
		{"empty string", "", 0},
		{"one segment", "abcdf", 1},
		{"testcase1", "hello! my way", 3},
		{"testcase2", "hello! mr lee's bag", 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := countSegments(tt.s)
			if out != tt.exp {
				t.Fatalf("with s:%s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
