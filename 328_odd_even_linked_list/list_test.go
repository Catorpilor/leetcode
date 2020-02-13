package list
import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)


func TestOddEvenList(t *testing.T) {
	st := []struct{
		name string
		head, exp *utils.ListNode
	}{
		{"empty list",nil, nil},
		{"single node", utils.ConstructFromSlice([]int{1}), utils.ConstructFromSlice([]int{1})},
		{"two nodes", utils.ConstructFromSlice([]int{1,2}), utils.ConstructFromSlice([]int{1,2})},
		{"testcase1",utils.ConstructFromSlice([]int{1,2,3,4,5}), utils.ConstructFromSlice([]int{1,3,5,2,4})},
	}
	for _, tt := range st {
		t.Run(tt.name, func(t *testing.T){
			out := oddEvenList(tt.head)
			if !utils.IsEqualList(out, tt.exp) {
				t.Fatalf("with head: %s wanted %s but got %s", tt.head.String(), tt.exp.String(), out.String() )
			}
			t.Log("pass")
		})
	}
}