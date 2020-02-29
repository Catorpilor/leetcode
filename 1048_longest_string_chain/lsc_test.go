package lsc

import "testing"

func TestLongesteStringChain(t *testing.T) {
	st := []struct {
		name  string
		words []string
		exp   int
	}{
		{"empty words", nil, 0},
		{"single words", []string{"a"}, 1},
		{"not connected", []string{"a", "bca"}, 1},
		{"testcase1", []string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestChain(tt.words)
			if out != tt.exp {
				t.Fatalf("with words: %v wanted %d but got %d", tt.words, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
