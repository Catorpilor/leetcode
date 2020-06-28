package pc

import "testing"

func TestPathCrossing(t *testing.T) {
	st := []struct {
		name string
		path string
		exp  bool
	}{
		{"testcase1", "NES", false},
		{"testcase2", "NESWW", true},
		{"testcase3", "NESW", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := pathCrossing(tt.path)
			if out != tt.exp {
				t.Fatalf("with input path %s, wanted %t but got %t", tt.path, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
