package msb

import "testing"

func TestMinWindow(t *testing.T) {
	st := []struct {
		name string
		s, t string
		exp  string
	}{
		{"empty s", "", "123", ""},
		{"empty t", "123", "", ""},
		{"s eq t", "ABC", "ABC", "ABC"},
		{"testcase1", "ADOBECODEBANC", "ABC", "BANC"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := MinWindow(tt.s, tt.t)
			if out != tt.exp {
				t.Fatalf("with input s:%s, t:%s, wanted %s but got %s", tt.s, tt.t, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
