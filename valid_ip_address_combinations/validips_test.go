package validips

import (
	"reflect"
	"testing"
)

func TestValidIPs(t *testing.T) {
	st := []struct {
		name string
		str  string
		exp  []string
	}{
		{"empty str", "", nil},
		{"invalid1", "123", nil},
		{"testcase1", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := validIPs(tt.str)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input str:%s wanted %v but got %v", tt.str, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
