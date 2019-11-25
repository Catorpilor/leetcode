package oed

import "testing"

func TestIsOneEditDistance(t *testing.T) {
	st := []struct {
		name         string
		word1, word2 string
		exp          bool
	}{
		{"both empty", "", "", false},
		{"s is empty", "", "123", false},
		{"s is empty, len(t) is 1", "", "1", true},
		{"both equal", "123", "123", false},
		{"testcase1", "123", "12", true},
		{"testcase2", "cab", "ad", false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isOneEditDistance(tt.word1, tt.word2)
			if out != tt.exp {
				t.Fatalf("with word1: %s and word2: %s wanted %t but got %t", tt.word1, tt.word2, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
