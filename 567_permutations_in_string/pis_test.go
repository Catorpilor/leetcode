package pis

import "testing"

func TestCheckInclusion(t *testing.T) {
	st := []struct {
		name   string
		s1, s2 string
		exp    bool
	}{
		{"empty s1", "", "abc", false},
		{"empty s2", "ab", "", false},
		{"s1 is longer than s2", "abcdfd", "abc", false},
		{"s1 = s2", "abc", "bca", true},
		{"testcase1", "ab", "eidbaooo", true},
		{"testcase2", "ab", "eidboaoo", false},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := checkInclusion(tt.s1, tt.s2)
			if out != tt.exp {
				t.Fatalf("with input s1: %s, s2: %s wanted %t but got %t", tt.s1, tt.s2,
					tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
