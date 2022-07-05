package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

func computeDepthAndBalanced(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	leftOk, left := computeDepthAndBalanced(root.Left)
	if !leftOk {
		return false, 0
	}

	rightOk, right := computeDepthAndBalanced(root.Right)
	if !rightOk {
		return false, 0
	}

	diff := left - right
	switch diff {
	case 1, 0:
		return true, left + 1
	case -1:
		return true, right + 1
	default:
		return false, 0
	}
}

func isBalanced(root *TreeNode) bool {
	isOk, _ := computeDepthAndBalanced(root)
	return isOk
}

func Test110() {
	fmt.Println(isBalanced(ToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})))
	fmt.Println(isBalanced(ToTreeNode([]int{1, 2, 2, 3, 3, 0, 0, 4, 4})) == false)
	fmt.Println(isBalanced(ToTreeNode([]int{})))
}

// @lc code=end
