package bulb

import "testing"


func TestMinFlips(t *testing.T){
	st := []struct{
		name string
		target string
		exp int
	}{
		{"testcase1", "10111", 3},
		{"testcase2", "000", 0},
		{"testcase3", "101", 3},
		{"testcase4", "001011101", 5},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			out := minFlips(tt.target)
			if out != tt.exp {
				t.Fatalf("with input target:%s wanted %d but got %d", tt.target, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
