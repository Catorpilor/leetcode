package wp

import "testing"

func TestIsMatch(t *testing.T) {
	st := []struct {
		name         string
		pattern, str string
		exp          bool
	}{
		{"both empty", "", "", true},
		{"str is empty", "ab", "", false},
		{"pattern is empty", "", "abc", false},
		{"testcase1", "abab", "redblueredblue", true},
		{"testcase2", "aaaa", "asdasdasdasd", true},
		{"testcase3", "aabb", "xyzabcxzyabc", false},
		{"failed1", "ab", "aa", false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isMatch(tt.pattern, tt.str)
			if out != tt.exp {
				t.Fatalf("with pattern:%s and str:%s wanted %t but got %t", tt.pattern, tt.str, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
