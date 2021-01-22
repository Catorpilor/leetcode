package ds

import "testing"

func TestIsClose(t *testing.T) {
	st := []struct {
		name         string
		word1, word2 string
		exp          bool
	}{
		{"both empty", "", "", true},
		{"equal", "abc", "abc", true},
		{"testcase1", "abc", "bca", true},
		{"testcase2", "a", "aa", false},
		{"testcase3", "cabbba", "abbccc", true},
		{"testcase4", "cabbba", "aabbss", false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isClose(tt.word1, tt.word2)
			if out != tt.exp {
				t.Fatalf("with input wanted %t but got %t", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
