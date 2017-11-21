package ranger

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	st := []struct {
		name string
		nums []int
		exp  []string
	}{
		{"empty slice", []int{}, []string{}},
		{"single", []int{1}, []string{"1"}},
		{"test 1", []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{"test 2", []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}
	for _, c := range st {
		t.Run(c.name, func(t *testing.T) {
			ret := SummaryRanges(c.nums)
			if !reflect.DeepEqual(ret, c.exp) {
				t.Fatalf("expected %v but got %v with input %v",
					c.exp, ret, c.nums)
			}
		})
	}

}
