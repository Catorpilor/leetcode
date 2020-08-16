package odds

import "testing"

func TestIsSatisfy(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  bool
	}{
		{"empty slice", nil, false},
		{"all odds", []int{1, 3, 5, 7, 9}, true},
		{"testcase1", []int{2, 6, 4, 1}, false},
		{"testcase2", []int{1, 2, 34, 3, 4, 5, 7, 9, 8}, true},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := isStatisfy(tt.nums)
			if out != tt.exp {
				t.Fatalf("with input nums:%v wanted %t but got %t", tt.nums, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
