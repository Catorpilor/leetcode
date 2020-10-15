package bst

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/catorpilor/leetcode/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

// Serializes a tree to a single string. we use the level order traversal here. time complexity O(N), space complexity O(N)
func (this *Codec) serialize(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}
	// just a level order traversal
	var vals []*utils.TreeNode
	vals = append(vals, root)
	var sb bytes.Buffer
	for len(vals) > 0 {
		top := vals[0]
		vals = vals[1:]
		sb.WriteString(strconv.Itoa(top.Val))
		if top.Left != nil {
			vals = append(vals, top.Left)
		}
		if top.Right != nil {
			vals = append(vals, top.Right)
		}
		if len(vals) > 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

// deserialize your encoded data to tree.
func (this *Codec) deserialize(data string) *utils.TreeNode {
	// construct a bst from level order traversal
	// split data by space
	sd := strings.Fields(data)
	var root *utils.TreeNode
	for i := range sd {
		val, _ := strconv.Atoi(string(sd[i]))
		root = levelOrder(root, val)
	}
	return root
}

func levelOrder(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		root = &utils.TreeNode{Val: val}
		return root
	}
	if val <= root.Val {
		root.Left = levelOrder(root.Left, val)
	} else {
		root.Right = levelOrder(root.Right, val)
	}
	return root

}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
