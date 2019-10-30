package utils

import (
	"testing"
)

func TestPow(t *testing.T){
	st := []struct{
		name string
		base, expr int
		exp int
	}{
		{"exp eq 1", 2, 1, 2},
		{"base eq 1", 1, 3, 1},
		{"test1", 2, 1, 2},
		{"test2", 2,5, 32},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			out := Pow(tt.base, tt.expr)
			if out != tt.exp {
				t.Fatalf("with input base: %d and expr: %d wanted %d but got %d", tt.base, tt.expr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
