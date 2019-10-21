package pp

import (
	"reflect"
	"testing"
)

func TestGeneratePalindromes(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []string
	}{
		{"empty string", "", []string{}},
		{"len eq 1", "a", []string{"a"}},
		{"two as", "aa", []string{"aa"}},
		{"test1", "abc", []string{}},
		{"test2", "aabb", []string{"abba", "baab"}},
		{"test3", "aba", []string{"aba"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := generatePalindromes(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s: %s wanted %v but got %v", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
