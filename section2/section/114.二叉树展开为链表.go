package section

import "fmt"

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flattenTreeNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	rightNode := flattenTreeNode(root.Right)
	if root.Left != nil {
		root.Right = flattenTreeNode(root.Left)
		root.Left = nil
		current := root.Right
		for current.Right != nil {
			current = current.Right
		}

		current.Right = rightNode
	}

	return root
}

func flatten(root *TreeNode) {
	flattenTreeNode(root)
}

func Test114() {
	nodes := ToTreeNode([]int{1, 2, 5, 3, 4, 0, 6})
	flatten(nodes)
	// [1,null,2,null,3,null,4,null,5,null,6]
	fmt.Println(nodes.ToValues())
}

// @lc code=end
