package robot

import "testing"

func TestIsCircle(t *testing.T) {
	st := []struct {
		name         string
		instructions string
		exp          bool
	}{
		{"gg", "GG", false},
		{"gl", "GL", true},
		{"testcase1", "GGLLGG", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isCircle(tt.instructions)
			if out != tt.exp {
				t.Fatalf("with input instructions:%v wanted %t  but got %t", tt.instructions, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
