package utils

import "testing"

func TestReverseList(t *testing.T) {
	st := []struct {
		name      string
		head, exp *ListNode
	}{
		{"empty list", nil, nil},
		{"single node", ConstructFromSlice([]int{1}), ConstructFromSlice([]int{1})},
		{"testcase1", ConstructFromSlice([]int{1, 2, 3}), ConstructFromSlice([]int{3, 2, 1})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := ReverseList(tt.head)
			if !IsEqualList(tt.exp, out) {
				t.Fatalf("with input list: %s wanted %s but got %s", tt.head, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
