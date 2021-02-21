package roman

import "testing"

func TestRomanToInt(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  int
	}{
		{"empty string", "", 0},
		{"testcase1", "III", 3},
		{"testcase2", "IX", 9},
		{"testcase3", "LVIII", 58},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {

			out := romanToInt(tt.s)
			if out != tt.exp {
				t.Fatalf("wanted %d but got %d", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
