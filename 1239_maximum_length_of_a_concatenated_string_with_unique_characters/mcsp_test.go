package mcsp

import "testing"

func TestMaxLength(t *testing.T) {
	st := []struct {
		name string
		arr  []string
		exp  int
	}{
		{"empty arr", []string{}, 0},
		{"allunique", []string{"abcdefg"}, 7},
		{"testcase1", []string{"un", "iq", "ue"}, 4},
		{"testcase2", []string{"cha", "r", "act", "ers"}, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxLength(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestMaxLengthWithDfs(t *testing.T) {
	st := []struct {
		name string
		arr  []string
		exp  int
	}{
		{"empty arr", []string{}, 0},
		{"allunique", []string{"abcdefg"}, 7},
		{"testcase1", []string{"un", "iq", "ue"}, 4},
		{"testcase2", []string{"cha", "r", "act", "ers"}, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxLength(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestMaxLengthWithBackTrack(t *testing.T) {
	st := []struct {
		name string
		arr  []string
		exp  int
	}{
		{"empty arr", []string{}, 0},
		{"allunique", []string{"abcdefg"}, 7},
		{"testcase1", []string{"un", "iq", "ue"}, 4},
		{"testcase2", []string{"cha", "r", "act", "ers"}, 6},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := maxLength(tt.arr)
			if out != tt.exp {
				t.Fatalf("with input arr:%v wanted %d but got %d", tt.arr, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
