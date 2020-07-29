package file

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetFolderNames(t *testing.T) {
	st := []struct {
		name  string
		names []string
		exp   []string
	}{
		{"empty names", nil, nil},
		{"testcase1", []string{"pes", "fifa", "gta", "pes(2019)"}, []string{"pes", "fifa", "gta", "pes(2019)"}},
		{"testcase2", []string{"gta", "gta(1)", "gta", "avalon"}, []string{"gta", "gta(1)", "gta(2)", "avalon"}},
		{"testcase3", []string{"op", "op(1)", "op(2)", "op(3)", "op"}, []string{"op", "op(1)", "op(2)", "op(3)", "op(4)"}},
		{"testcase4", []string{"kd", "kd(1)", "kd", "kd(1)"}, []string{"kd", "kd(1)", "kd(2)", "kd(1)(1)"}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := getFolderNames(tt.names)
			if diff := cmp.Diff(tt.exp, out); diff != "" {
				t.Fatalf("with input folders:%v (+want, -got) %s", tt.names, diff)
			}
			t.Log("pass")
		})
	}
}
