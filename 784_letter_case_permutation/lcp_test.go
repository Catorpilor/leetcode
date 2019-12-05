package lcp

import (
	"reflect"
	"testing"
)

func TestLetterCasePermutations(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []string
	}{
		{"empty", "", []string{}},
		{"all numeric", "123", []string{"123"}},
		{"testcase1", "a1b2", []string{"a1B2", "A1b2", "a1b2", "A1B2"}},
		{"testcase2", "3z4", []string{"3z4", "3Z4"}},
		{"failed1", "C", []string{"c", "C"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := letterCasePermutations(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s:%s wanted %v but got %v", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
