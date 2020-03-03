package cc

import "testing"

func TestCriticalConnections(t *testing.T) {
	st := []struct {
		name string
		n    int
		conn [][]int
		exp  [][]int
	}{
		{"n = 1", 1, nil, nil},
		{"n = 2", 2, [][]int{[]int{1, 0}}, [][]int{[]int{1, 0}}},
		{"testcase1", 4, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 0}, []int{1, 3}}, [][]int{[]int{1, 3}}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := criticalConns(tt.n, tt.conn)
			if len(out) != len(tt.exp) {
				t.Fatalf("with n: %d and conn: %v wanted %d but got %d critical connections", tt.n, tt.conn, len(tt.exp), len(out))
			}
		})
	}
}
