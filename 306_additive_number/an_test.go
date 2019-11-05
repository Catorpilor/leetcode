package an

import "testing"

func TestIsAdditive(t *testing.T) {
	st := []struct {
		name string
		num  string
		exp  bool
	}{}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isAdditive(tt.num)
			if out != tt.exp {
				t.Fatalf("with input num:%s wanted %t but got %t", tt.num, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
