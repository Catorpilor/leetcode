package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	low, high := 0, len(nums)-1
	mid := (high-low)/2 + low
	root := &TreeNode{Val: nums[mid]}
	// create left sub tree
	root.Left = sortedArrayToBST(nums[:mid])
	// create right sub tree
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
