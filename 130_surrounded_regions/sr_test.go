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
		{"all xes", [][]byte{[]byte{'x', 'x'}, []byte{'x', 'x'}}, [][]byte{[]byte{'x', 'x'}, []byte{'x', 'x'}}},
		{"all oes", [][]byte{[]byte{'o', 'o', 'o'}, []byte{'o', 'o', 'o'}, []byte{'o', 'o', 'o'}}, [][]byte{[]byte{'o', 'o', 'o'}, []byte{'o', 'x', 'o'}, []byte{'o', 'o', 'o'}}},
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
