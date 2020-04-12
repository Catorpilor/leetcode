package candy

import "testing"

func TestNumOfCandies(t *testing.T) {
	st := []struct {
		name    string
		ratings []int
		exp     int
	}{
		{"empty list", nil, 0},
		{"single one", []int{1}, 1},
		{"all identical", []int{1, 1, 1, 1}, 4},
		{"testcase1", []int{1, 0, 2}, 5},
		{"testcase2", []int{1, 2, 2}, 4},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := numOfCandies(tt.ratings)
			if out != tt.exp {
				t.Fatalf("with input ratings:%v wanted %d but got %d", tt.ratings, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
