package matrix

import "testing"

func TestMaxMatrix(t *testing.T) {
	st := []struct {
		name   string
		matrix [][]byte
		exp    int
	}{
		{"empty", [][]byte{}, 0},
		{"single 1", [][]byte{[]byte{'1'}}, 1},
		{"test 1", [][]byte{[]byte{'1', '0'}, []byte{'0', '1'}}, 1},
		{"test 2", [][]byte{[]byte{'1', '1'}, []byte{'1', '0'}}, 2},
		{"test 3", [][]byte{[]byte{'1', '0', '1', '0', '0'}, []byte{'1', '0', '1', '1', '1'},
			[]byte{'1', '1', '1', '1', '1'}, []byte{'1', '0', '0', '1', '0'}}, 6},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxMatrix(c.matrix)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.matrix)
			}
		})
	}
}
