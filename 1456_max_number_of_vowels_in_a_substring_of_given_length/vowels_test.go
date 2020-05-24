package vowels

import "testing"

func TestMaxVowels(t *testing.T) {
	st := []struct {
		name string
		s    string
		k    int
		exp  int
	}{
		{"k eq 1", "abc", 1, 1},
		{"k eq len(s)", "abc", 3, 1},
		{"testcase1", "abciiidef", 3, 3},
		{"testcase2", "aeiou", 2, 2},
		{"testcase3", "leetcode", 3, 2},
		{"testcase4", "rhythms", 4, 0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxVowels(tt.s, tt.k)
			if out != tt.exp {
				t.Fatalf("with input s:%s and k:%d wanted %d but got %d", tt.s, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
