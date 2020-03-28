package interleave

import "testing"

func TestIsInterLeaving(t *testing.T) {
	st := []struct {
		name       string
		s1, s2, s3 string
		exp        bool
	}{
		{"all empty", "", "", "", true},
		{"s1 is empty", "", "abc", "abc", true},
		{"len(s1+s2) != len(s3)", "abc", "cde", "abcde", false},
		{"testcase1", "aabcc", "dbbca", "aadbbcbcac", true},
		{"failed1", "ab", "bc", "babc", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isInterLeaving(tt.s1, tt.s2, tt.s3)
			if out != tt.exp {
				t.Fatalf("with input s1:%s, s2:%s and s3:%s wanted %t but got %t", tt.s1, tt.s2, tt.s3, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
