package caps

import "testing"

func TestDetectCaps(t *testing.T) {
	st := []struct {
		name string
		word string
		exp  bool
	}{
		{"all caps", "CHINA", true},
		{"all lowercase", "leetcode", true},
		{"testcase1", "Flag", true},
		{"testcase2", "FlaG", false},
		{"testcase3", "leetCode", false},
		{"testcase4", "FFFFFFf", false},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := detectCaps(tt.word)
			if out != tt.exp {
				t.Fatalf("with input word:%s wanted %t but got %t", tt.word, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
