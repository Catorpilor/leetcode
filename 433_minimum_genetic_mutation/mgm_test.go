package mgm

import "testing"

func TestMinMutation(t *testing.T) {
	st := []struct {
		name       string
		start, end string
		bank       []string
		exp        int
	}{
		{"empty bank", "AACC", "AATT", []string{}, -1},
		{"testcase1", "AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"testcase2", "AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"failed1", "AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}, -1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := minMutation(tt.start, tt.end, tt.bank)
			if out != tt.exp {
				t.Fatalf("with start:%s and end:%s and bank:%v wanted %d but got  %d", tt.start, tt.end, tt.bank, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
