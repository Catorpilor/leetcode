package dots

import "testing"

func TestMinDistance(t *testing.T) {
	st := []struct {
		name         string
		word1, word2 string
		exp          int
	}{
		{"both empty", "", "", 0},
		{"word1 is empty", "", "abc", 3},
		{"word1 eq word2", "abc", "abc", 0},
		{"testcase1", "sea", "eat", 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minDistance(tt.word1, tt.word2)
			if out != tt.exp {
				t.Fatalf("with word1:%s and word2: %s, wanted %d but got %d", tt.word1, tt.word2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
