package ld

import "testing"

func TestLongestWord(t *testing.T) {
	st := []struct {
		name string
		s    string
		dict []string
		exp  string
	}{
		{"testcase1", "abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"testcase2", "abpcplea", []string{"a", "b", "c"}, "a"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestWord(tt.s, tt.dict)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
