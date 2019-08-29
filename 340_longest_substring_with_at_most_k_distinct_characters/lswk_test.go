package lswk

import "testing"

func TestLongestSubstringAtmostK(t *testing.T) {
	st := []struct {
		name string
		s    string
		k    int
		exp  int
	}{
		{"empty string", "", 1, 0},
		{"invalid k=0", "abc", 0, 0},
		{"invalid k<0", "abc", -1, 0},
		{"k is larger than len(s)", "abcdfs", 7, 6},
		{"testcase1", "eceba", 2, 3},
		{"testcase2", "aa", 1, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := LongestSubstringAtmostK(tt.s, tt.k)
			if out != tt.exp {
				t.Fatalf("with input s:%s, k:%d, wanted %d but got %d", tt.s, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
