package lcs

import "testing"

func TestLCS(t *testing.T) {
	st := []struct {
		name         string
		text1, text2 string
		exp          int
	}{
		{"both empty", "", "", 0},
		{"text1 is empty", "", "acb", 0},
		{"text1 eq text2", "abc", "abc", 3},
		{"testcase1", "abcde", "ace", 3},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := longestCommonSeq(tt.text1, tt.text2)
			if out != tt.exp {
				t.Fatalf("with text1: %s and text2: %s wanted %d but got %d", tt.text1, tt.text2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
