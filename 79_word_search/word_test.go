package word

import "testing"

func TestExists(t *testing.T) {
	st := []struct {
		name  string
		board [][]byte
		word  string
		exp   bool
	}{
		{"nil board", nil, "test", false},
		{"1*3 board", [][]byte{[]byte{'a', 'b', 'c'}}, "abcd", false},
		{"test 1", [][]byte{[]byte{'a', 'b', 'c', 'e'}, []byte{'s', 'f', 'c', 's'},
			[]byte{'a', 'd', 'e', 'e'}}, "see", true},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := Exists(c.board, c.word)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t with input %v and %s",
					c.exp, ret, c.board, c.word)
			}
		})
	}

}
