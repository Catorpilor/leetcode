package kd

import "testing"

func TestRemoveDigits(t *testing.T) {
	st := []struct {
		name string
		num  string
		k    int
		exp  string
	}{
		{"empty string", "", 1, ""},
		{"k eq to len(s)", "123", 3, "0"},
		{"testcase1", "1423319", 3, "1219"},
		{"testcase2", "10200", 1, "200"},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := removeKDigits(tt.num, tt.k)
			if out != tt.exp {
				t.Fatalf("with input num:%s and k:%d wanted %s but got %s", tt.num, tt.k, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
