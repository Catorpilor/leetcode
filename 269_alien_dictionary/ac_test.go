package ac

import "testing"

func TestAlienOrder(t *testing.T) {
	st := []struct {
		name  string
		words []string
		exp   string
	}{
		{"empty words", []string{}, ""},
		{"test1", []string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{"test2", []string{"z", "x"}, "zx"},
		{"test3", []string{"z", "x", "z"}, ""},
		{"failed1", []string{"za", "zb", "ca", "cb"}, "abzc"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := alienOrder(tt.words)
			if out != tt.exp {
				t.Fatalf("with input words: %v wanted %s but got %s", tt.words, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
