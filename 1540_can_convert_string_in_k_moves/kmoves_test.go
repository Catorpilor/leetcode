package kmoves

import "testing"

func TestCanConvert(t *testing.T) {
	st := []struct {
		name string
		s, t string
		k    int
		exp  bool
	}{
		{"len(s) != len(t)", "abc", "defg", 1, false},
		{"identical", "abc", "abc", 2, true},
		{"testcase1", "input", "ouput", 9, true},
		{"testcase2", "abc", "bcd", 10, false},
		{"testcase3", "aab", "bbb", 27, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := canConvert(tt.s, tt.t, tt.k)
			if out != tt.exp {
				t.Fatalf("with input s:%s and t:%s and k:%d wanted %t but got %t", tt.s, tt.t, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
