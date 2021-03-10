package roman

import "testing"

func TestIntToRoman(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  string
	}{
		{"testcase1", 1, "I"},
		{"testcase2", 4, "IV"},
		{"testcase3", 3, "III"},
		{"testcase4", 58, "LVIII"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := intToRoman(tt.num)
			if out != tt.exp {
				t.Fatalf("wanted %s but got %s", tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
