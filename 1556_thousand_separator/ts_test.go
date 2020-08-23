package ts

import "testing"

func TestThousandSep(t *testing.T) {
	st := []struct {
		name string
		num  int
		exp  string
	}{
		{"num=0", 0, "0"},
		{"testcase1", 123, "123"},
		{"testcase2", 1234, "1.234"},
		{"testcase3", 1000000, "1.000.000"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := addSeps(tt.num)
			if out != tt.exp {
				t.Fatalf("with input num:%d wanted %s but got %s", tt.num, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
