package anagrams

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	st := []struct {
		name, s, p string
		exp        []int
	}{
		{"non exits", "abab", "cd", nil},
		{"example 1", "cbaebabacd", "abc", []int{0, 6}},
		{"example 2", "abab", "ab", []int{0, 1, 2}},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := FindAnagrams(c.s, c.p)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v, with input %s and %s",
					c.exp, ret, c.s, c.p)
			}
		})
	}
}
