package ws

import (
	"reflect"
	"testing"
)

func TestWordSquares(t *testing.T) {
	st := []struct {
		name  string
		words []string
		exp   [][]string
	}{
		{"empty", []string{}, [][]string{}},
		{"testcase1", []string{"area", "lead", "wall", "lady", "ball"}, [][]string{[]string{"wall", "area", "lead", "lady"},
			[]string{"ball", "area", "lead", "lady"}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := wordSquares(tt.words)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with words: %v wanted %v but got %v", tt.words, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
