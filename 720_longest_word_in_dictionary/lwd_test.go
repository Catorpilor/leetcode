package lwd

import "testing"

func TestLongestWord(t *testing.T) {
	st := []struct {
		name  string
		words []string
		exp   string
	}{
		{"nil words", nil, ""},
		{"empty words", []string{}, ""},
		{"single words", []string{"hello"}, "hello"},
		{"test 1", []string{"w", "wo", "wor", "worl", "world"}, "world"},
		{"test 2", []string{"a", "banana", "app", "ap", "apply", "appl", "apple"}, "apple"},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := LongestWord(c.words)
			if ret != c.exp {
				t.Fatalf("expected %d but got %s, with input %v",
					c.exp, ret, c.words)
			}
		})
	}
}
