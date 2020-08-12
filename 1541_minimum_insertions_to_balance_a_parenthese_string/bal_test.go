package bal

import "testing"

func TestMinInserts(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"single parenthe", "(", 2},
		{"testcase1", "(()))", 1},
		{"testcase2", "())", 0},
		{"testcase3", "((((((", 12},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minInserts(tt.s)
			if out != tt.exp {
				t.Fatalf("with input s:%s wanted %d but got %d", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
