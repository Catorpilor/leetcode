package tower

import "testing"

func TestChampangneTower(t *testing.T) {
	st := []struct {
		name                           string
		poured, query_row, query_glass int
		exp                            float64
	}{
		{"testcase1", 1, 1, 1, 0.0},
		{"testcase2", 2, 1, 1, 0.5},
		{"testcase3", 100000009, 33, 17, 1.0},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := queryGlass(tt.poured, tt.query_row, tt.query_glass)
			if out != tt.exp {
				t.Fatalf("with input (%d, %d, %d), wanted %.2f but got %.2f", tt.poured, tt.query_row, tt.query_glass, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
