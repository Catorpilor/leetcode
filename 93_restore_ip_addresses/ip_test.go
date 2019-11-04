package ip

import (
	"reflect"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []string
	}{
		{"empty str", "", []string{}},
		{"not a valid ip", "012313321", []string{}},
		{"testcase1", "25525511125", []string{"255.255.111.25", "255.255.11.125"}},
		{"testcase2", "2552552550", []string{"255.255.255.0", "255.255.25.50"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := restoreIpAddresses(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s:%s wanted %v but got %v", tt.s, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
