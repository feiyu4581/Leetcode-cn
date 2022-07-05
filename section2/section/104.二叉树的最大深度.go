package section

import "fmt"

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func recursionMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := recursionMaxDepth(root.Left)
	rightDepth := recursionMaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func maxDepth(root *TreeNode) int {
	return recursionMaxDepth(root)
}

func Test104() {
	fmt.Println(maxDepth(ToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})) == 3)
}

// @lc code=end
