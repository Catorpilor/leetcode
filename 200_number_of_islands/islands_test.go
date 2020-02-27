package islands

import "testing"

func TestNumOfIslands(t *testing.T) {
	st := []struct {
		name string
		grid [][]byte
		exp  int
	}{
		{"empty", [][]byte{}, 0},
		{"single", [][]byte{[]byte{'1'}}, 1},
		{"single0", [][]byte{[]byte{'0'}}, 0},
		{"test1", [][]byte{[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '1'}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := numOfIslands(c.grid)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.grid)
			}
		})
	}
}

func TestUseBFS(t *testing.T) {
	st := []struct {
		name string
		grid [][]byte
		exp  int
	}{
		{"empty", [][]byte{}, 0},
		{"single", [][]byte{[]byte{'1'}}, 1},
		{"single0", [][]byte{[]byte{'0'}}, 0},
		{"test1", [][]byte{[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '1'}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useBfs(c.grid)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.grid)
			}
		})
	}
}

func TestUseUnionFind(t *testing.T) {
	st := []struct {
		name string
		grid [][]byte
		exp  int
	}{
		{"empty", [][]byte{}, 0},
		{"single", [][]byte{[]byte{'1'}}, 1},
		{"single0", [][]byte{[]byte{'0'}}, 0},
		{"test1", [][]byte{[]byte{'1', '1', '1', '1', '0'},
			[]byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'},
			[]byte{'0', '0', '0', '0', '1'}}, 2},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := useUnionFind(c.grid)
			if ret != c.exp {
				t.Fatalf("expected %d but got %d with input %v",
					c.exp, ret, c.grid)
			}
		})
	}
}
