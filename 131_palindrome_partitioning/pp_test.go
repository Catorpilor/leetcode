package pp

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  [][]string
	}{
		{"empty", "", [][]string{}},
		{"single letter", "a", [][]string{[]string{"a"}}},
		{"3 letters identical", "aaa", [][]string{[]string{"a", "a", "a"}, []string{"a", "aa"}, []string{"aa", "a"}, []string{"aaa"}}},
		{"unique string", "abc", [][]string{[]string{"a", "b", "c"}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := partition(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s: %s wanted %v but got %v", tt.s, tt.exp, out)
			}
			t.Log("psss")
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  bool
	}{
		{"empty", "", true},
		{"single", "a", true},
		{"testcase1", "aa", true},
		{"testcase2", "abc", false},
		{"testcase3", "aba", true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isPalindrome(tt.s)
			if tt.exp != out {
				t.Fatalf("with input s: %s wanted %t but got %t", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
