package bomb

import "testing"

func TestMaxKill(t *testing.T) {
	st := []struct {
		name   string
		grid   [][]byte
		expect int
	}{
		{"nil grid", nil, 0},
		{"empty grid", [][]byte{}, 0},
		{"test 1", [][]byte{[]byte{'0', 'E', '0', '0'}, []byte{'E', '0', 'W', 'E'}, []byte{'0', 'E', '0', '0'}}, 3},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := MaxKill(c.grid)
			if ret != c.expect {
				t.Fatalf("expected %d but got %d, with input %v",
					c.expect, ret, c.grid)
			}
		})
	}
}
