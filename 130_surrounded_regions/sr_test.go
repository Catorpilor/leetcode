package sr

import (
	"reflect"
	"testing"
)

func TestRegions(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		exp   [][]byte
	}{
		{"empty board", [][]byte{}, [][]byte{}},
		{"all xes", [][]byte{[]byte{'X', 'X'}, []byte{'X', 'X'}}, [][]byte{[]byte{'X', 'X'}, []byte{'X', 'X'}}},
		{"all oes", [][]byte{[]byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}}, [][]byte{[]byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := solved(tt.board)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input board: %v wanted %v but got %v", tt.board, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestBfs(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		exp   [][]byte
	}{
		{"empty board", [][]byte{}, [][]byte{}},
		{"all xes", [][]byte{[]byte{'X', 'X'}, []byte{'X', 'X'}}, [][]byte{[]byte{'X', 'X'}, []byte{'X', 'X'}}},
		{"all oes", [][]byte{[]byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}}, [][]byte{[]byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}, []byte{'O', 'O', 'O'}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := bfs(tt.board)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input board: %v wanted %v but got %v", tt.board, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
