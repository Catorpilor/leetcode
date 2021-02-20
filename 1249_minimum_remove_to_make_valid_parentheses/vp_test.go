package vp

import "testing"

func TestMinRemove(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  string
	}{
		{"testcase1", "lee(t(c)o)de)", "lee(t(co)de)"},
		{"testcase2", "a)b(c)d", "ab(c)d"},
		{"testcase3", "))((", ""},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minRemove(tt.s)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
