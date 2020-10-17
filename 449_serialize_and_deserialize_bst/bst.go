package bst

import (
	"bytes"
	"math"
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

func (this *Codec) serializeUseDFS(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}
	var sb bytes.Buffer
	helper(&sb, root)
	return sb.String()
}

func helper(sb *bytes.Buffer, root *utils.TreeNode) {
	if root == nil {
		return
	}
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(",")
	helper(sb, root.Left)
	helper(sb, root.Right)
}

func (this *Codec) deserializeUseDFS(data string) *utils.TreeNode {
	if len(data) < 1 {
		return nil
	}
	sd := strings.FieldsFunc(data, func(r rune) bool {
		return r == ','
	})
	return insert(sd, math.MinInt32, math.MaxInt32)
}

func insert(sd []string, min, max int) *utils.TreeNode {
	if len(sd) == 0 {
		return nil
	}
	h := sd[0]
	ih, _ := strconv.Atoi(h)
	if ih < min || ih > max {
		return nil
	}
	sd = sd[1:]
	root := &utils.TreeNode{Val: ih}
	root.Left = insert(sd, min, ih)
	root.Right = insert(sd, ih, max)
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
