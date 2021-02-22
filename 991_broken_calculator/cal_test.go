package cal

import "testing"

func TestBrokenCalc(t *testing.T) {
	st := []struct {
		name string
		x, y int
		exp  int
	}{
		{"x equals y", 2, 2, 0},
		{"x > y", 1024, 1, 1023},
		{"testcase1", 3, 10, 3},
		{"testcase2", 2, 3, 2},
		{"testcase3", 5, 8, 2},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := brokenCalc(tt.x, tt.y)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")

		})
	}

}
