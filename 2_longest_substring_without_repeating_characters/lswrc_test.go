package lswrc

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	st := []struct {
		name, str string
		exp       int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"duplicate charater", "aaaa", 1},
		{"abcabcbb", "abcabcbb", 3},
		{"pwwkew", "pwwkew", 3},
	}
	for _, s := range st {
		t.Run(s.name, func(t *testing.T) {
			ret := lengthOfLongestSubstring(s.str)
			if ret != s.exp {
				t.Fatalf("expected %d but got %d with input string: %s",
					s.exp, ret, s.str)
			}
		})
	}
}

func TestBruteForceLengthOfLongestSubstring(t *testing.T) {
	st := []struct {
		name, str string
		exp       int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"duplicate charater", "aaaa", 1},
		{"abcabcbb", "abcabcbb", 3},
		{"pwwkew", "pwwkew", 3},
	}
	for _, s := range st {
		t.Run(s.name, func(t *testing.T) {
			ret := bruteForceLength(s.str)
			if ret != s.exp {
				t.Fatalf("expected %d but got %d with input string: %s",
					s.exp, ret, s.str)
			}
		})
	}
}

func TestLengthOfLongestSubstringOpt(t *testing.T) {
	st := []struct {
		name, str string
		exp       int
	}{
		{"empty string", "", 0},
		{"single character", "a", 1},
		{"duplicate charater", "aaaa", 1},
		{"abcabcbb", "abcabcbb", 3},
		{"pwwkew", "pwwkew", 3},
	}
	for _, s := range st {
		t.Run(s.name, func(t *testing.T) {
			ret := lengthOfLongestSubstringOpt(s.str)
			if ret != s.exp {
				t.Fatalf("expected %d but got %d with input string: %s",
					s.exp, ret, s.str)
			}
		})
	}
}
