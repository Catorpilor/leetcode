package deepCopy

import (
	"math"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	st := []struct {
		name      string
		head, exp *RNode
	}{
		{"head is nil", nil, nil},
		{"single node", constructFromSlice([][2]int{[2]int{1, math.MinInt32}}), constructFromSlice([][2]int{[2]int{1, math.MinInt32}})},
		{"testcase1", constructFromSlice([][2]int{[2]int{7, math.MinInt32}, [2]int{13, 0}, [2]int{11, 4}, [2]int{10, 2}, [2]int{1, 0}}), constructFromSlice([][2]int{[2]int{7, math.MinInt32}, [2]int{13, 0}, [2]int{11, 4}, [2]int{10, 2}, [2]int{1, 0}})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := dcopy(tt.head)
			if !isEqualRNode(out, tt.exp) {
				t.Fatalf("with head: %v wanted %v but got %v", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
