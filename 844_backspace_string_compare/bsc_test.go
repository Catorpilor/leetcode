package bsc

import "testing"

func TestBackspaceCompare(t *testing.T) {
	st := []struct {
		name string
		s, t string
		exp  bool
	}{
		{"both empty", "", "", true},
		{"all #s", "###", "###", true},
		{"testcase1", "ab#", "b#a", true},
		{"testcase2", "a##c", "#a#c", true},
		{"testcase3", "a#c", "b", false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := backspaceCompare(tt.s, tt.t)
			if out != tt.exp {
				t.Fatalf("with input s:%s and t:%s wanted %t but got %t", tt.s, tt.t, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
