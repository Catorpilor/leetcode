package etf

import (
	"reflect"
	"testing"
)

func TestExlusiveTime(t *testing.T) {
	st := []struct {
		name string
		n    int
		logs []string
		exp  []int
	}{
		{"testcase1", 1, []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}, []int{8}},
		{"testcase2", 2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, []int{3, 4}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := exclusiveTime(tt.n, tt.logs)
			if !reflect.DeepEqual(out, tt.exp) {
				t.Fatalf("with input n:%d and logs: %v wanted %v but got %v", tt.n, tt.logs, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
