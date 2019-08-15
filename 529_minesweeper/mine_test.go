package mine

import (
	"reflect"
	"testing"
)

func TestUpdateBoard(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		click []int
		exp   [][]byte
	}{
		{"empty board", [][]byte{}, []int{1, 2}, [][]byte{}},
		{"invalid click", [][]byte{[]byte{'B', 'B', 'M', 'B'}, []byte{'B', 'B', 'B', 'B'}},
			[]int{1, 2, 3}, [][]byte{[]byte{'B', 'B', 'M', 'B'}, []byte{'B', 'B', 'B', 'B'}}},
		{"test case 1", [][]byte{[]byte{'E', 'E', 'E', 'E', 'E'}, []byte{'E', 'E', 'M', 'E', 'E'},
			[]byte{'E', 'E', 'E', 'E', 'E'}, []byte{'E', 'E', 'E', 'E', 'E'}}, []int{3, 0},
			[][]byte{[]byte{'B', '1', 'E', '1', 'B'}, []byte{'B', '1', 'M', '1', 'B'},
				[]byte{'B', '1', '1', '1', 'B'}, []byte{'B', 'B', 'B', 'B', 'B'}}},
	}

	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := UpdateBoard(tt.board, tt.click)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input board %v and click %v, wanted %v but got %v",
					tt.board, tt.click, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
