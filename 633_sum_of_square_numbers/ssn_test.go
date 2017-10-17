package ssn

import "testing"

func TestJudgeSSN(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  bool
	}{
		{"num eq 1", 1, true},
		{"num eq 2", 2, true},
		{"num eq 5", 5, true},
		{"num eq 6", 6, false},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := JudgeSSN(c.num)
			if ret != c.exp {
				t.Fatalf("expected %t but got %t, with input %d",
					c.exp, ret, c.num)
			}
		})
	}
}
