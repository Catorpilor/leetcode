package pascal

import "testing"
import "reflect"

func TestGetRow(t *testing.T) {
	cases := []struct {
		name     string
		rowIndex int
		expect   []int
	}{
		{"rowIndex is 0", 0, []int{1}},
		{"rowIndex is 3", 3, []int{1, 3, 3, 1}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := GetRow(c.rowIndex)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v, but got %v, with input %d",
					c.expect, ret, c.rowIndex)
			}
		})
	}
	t.Log("pass")
}
