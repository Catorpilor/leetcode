package record

import "testing"

func TestCheckRecords(t *testing.T) {
	st := []struct {
		name, s string
		exp     bool
	}{
		{"PPALLP", "PPALLP", true},
		{"PPALLL", "PPALLL", false},
		{"PALPLL", "PALPLL", true},
		{"PALPLLPLLPL", "PALPLLPLLPL", true},
	}

	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := CheckRecords(c.s)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t,with input %s",
					c.exp, ret, c.s)
			}
		})
	}
}
