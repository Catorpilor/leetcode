package reverse

import "testing"

func TestReverseBits(t *testing.T) {
	cases := []struct {
		name   string
		input  uint32
		expect uint32
	}{
		{"single digit", 1, 2147483648},
		{"0", 0, 0},
		{"multiple digits", 43261596, 964176192},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := ReverseBits(c.input)
			if ret != c.expect {
				t.Fatalf("expected %d but got %d, with input %d",
					c.expect, ret, c.input)
			}
		})
	}
	t.Log("pass")
}
