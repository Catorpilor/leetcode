package sarr

import "testing"

func TestIsEqualString(t *testing.T) {
	st := []struct {
		name         string
		word1, word2 []string
		exp          bool
	}{
		{"all empty", nil, nil, true},
		{"testcase1", []string{"a", "bc"}, []string{"ab", "c"}, true},
		{"testcase2", []string{"a", "cb"}, []string{"ab", "c"}, false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isEqual(tt.word1, tt.word2)
			if out != tt.exp {
				t.Fatalf("with input word1:%v and word2:%v wanted %t but got %t", tt.word1, tt.word2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
