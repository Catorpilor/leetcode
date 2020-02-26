package dbl

import (
	"testing"

	"github.com/catorpilor/leetcode/utils"
)

func TestInOrderRec(t *testing.T) {
	root := utils.ConstructTree([]int{4, 2, 5, 1, 3})
	if root.Val != 4 {
		t.Fatalf("root'Val should be 4 but got %d", root.Val)
	}
	head := useRecur(root)
	for i := 1; i < 6; i++ {
		if head.Val != i {
			t.Logf("wanted %d but got %d", i, head.Val)
		}
		head = head.Right
	}
	t.Log("pass")
}

func TestInOrderWithStack(t *testing.T) {
	root := utils.ConstructTree([]int{4, 2, 5, 1, 3})
	if root.Val != 4 {
		t.Fatalf("root'Val should be 4 but got %d", root.Val)
	}
	head := useStack(root)
	for i := 1; i < 6; i++ {
		if head.Val != i {
			t.Logf("wanted %d but got %d", i, head.Val)
		}
		head = head.Right
	}
	t.Log("pass")
}
