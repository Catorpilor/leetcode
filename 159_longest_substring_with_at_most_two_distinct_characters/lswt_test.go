package lswt

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty s", "", 0},
		{"len eq 1", "a", 1},
		{"identical characters", "aaaaa", 5},
		{"two identicals", "aaaabbbb", 8},
		{"no duplicates", "abcde", 2},
		{"testcase1", "eceba", 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := LengthOfLongestSubstring(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s, wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
