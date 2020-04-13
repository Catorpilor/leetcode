package stone

import "testing"

func TestLastWeight(t *testing.T) {
	st := []struct {
		name   string
		stones []int
		exp    int
	}{
		{"empty", nil, 0},
		{"single stone", []int{1}, 1},
		{"odds all identical", []int{1, 1, 1}, 1},
		{"even all identical", []int{1, 1, 1, 1}, 0},
		{"testcase1", []int{2, 7, 4, 1, 8, 1}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := lastWeightStones(tt.stones)
			if out != tt.exp {
				t.Fatalf("with input stones: %v wanted %d but got %d", tt.stones, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}

func TestLastWeightUseBucket(t *testing.T) {
	st := []struct {
		name   string
		stones []int
		exp    int
	}{
		{"empty", nil, 0},
		{"single stone", []int{1}, 1},
		{"odds all identical", []int{1, 1, 1}, 1},
		{"even all identical", []int{1, 1, 1, 1}, 0},
		{"testcase1", []int{2, 7, 4, 1, 8, 1}, 1},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := useBucket(tt.stones)
			if out != tt.exp {
				t.Fatalf("with input stones: %v wanted %d but got %d", tt.stones, tt.exp, out)
			}
			t.Log("pass")
		})
	}
}
