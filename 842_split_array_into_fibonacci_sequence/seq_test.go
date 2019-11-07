package seq

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestSplitIntoFib(t *testing.T) {
	st := []struct {
		name string
		s    string
		exp  []int
	}{
		{"len(s)=1", "2", []int{}},
		{"len(s)=2", "12", []int{}},
		{"valid len(s)=3", "123", []int{1, 2, 3}},
		{"invalid with leading 0", "011", []int{0, 1, 1}},
		{"valid with 0", "101", []int{1, 0, 1}},
		{"testcase1", "112358130", []int{}},
		{"testcase2", "11235813", []int{1, 1, 2, 3, 5, 8, 13}},
		{"failed1", "539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511", []int{}},
		{"faield2", `3611537383985343591834441270352104793375145479938855071433500231900737525076071514982402115895535257195564161509167334647108949738176284385285234123461518508746752631120827113919550237703163294909`, []int{}},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T) {
			out := splitIntoFib(tt.s)
			if !reflect.DeepEqual(tt.exp, out) {
				t.Fatalf("with input s:%s wanted %v but got %v", tt.s, tt.exp, out)
			}
			genStr := gen(out)
			if tt.s != genStr {
				t.Fatalf("origin str: %s and genrated str: %s", tt.s, genStr)
			}
			t.Log("pass")
		})
	}

}

func gen(fibs []int) string {
	var sb strings.Builder
	for _, v := range fibs {
		sb.WriteString(strconv.FormatInt(int64(v), 10))
	}
	return sb.String()
}
