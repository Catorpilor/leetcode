package nge

import "testing"

func TestNextGreaterElement(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  int
	}{
		{"zero", 0, -1},
		{"one digit", 1, -1},
		{"descresing", 321, -1},
		{"increasing", 132, 213},
		{"testcase1", 12, 21},
		{"testcase2", 21, -1},
		{"testcase3", 139, 193},
		{"failed1", 12222333, 12223233},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := nextGreaterElement(tt.num)
			if out != tt.exp {
				t.Fatalf("with input num: %d ,wanted %d but got %d", tt.num, tt.exp, out)
			}
			t.Log("pass")
		})
	}

}
