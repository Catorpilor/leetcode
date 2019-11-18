package ws

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		words []string
		exp   []string
	}{
		{"empty words", [][]byte{[]byte{'a', 'b'}}, []string{}, []string{}},
		{"testcase1", [][]byte{[]byte{'o', 'a', 'a', 'n'}, []byte{'e', 't', 'a', 'e'}, []byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{"failed1", [][]byte{[]byte{'a', 'a'}}, []string{"aa"}, []string{"aa"}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := findWords(tt.board, tt.words)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with boards: %v and words: %v wanted %v but got %v", tt.board, tt.words, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestWithTrie(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		words []string
		exp   []string
	}{
		{"empty words", [][]byte{[]byte{'a', 'b'}}, []string{}, []string{}},
		{"testcase1", [][]byte{[]byte{'o', 'a', 'a', 'n'}, []byte{'e', 't', 'a', 'e'}, []byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{"failed1", [][]byte{[]byte{'a', 'a'}}, []string{"aa"}, []string{"aa"}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := withTrie(tt.board, tt.words)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with boards: %v and words: %v wanted %v but got %v", tt.board, tt.words, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
