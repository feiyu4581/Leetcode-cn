package section

import "fmt"

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func Test112() {
	fmt.Println(hasPathSum(ToTreeNode([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}), 22))
	fmt.Println(hasPathSum(ToTreeNode([]int{-2, 0, -3}), -5) == true)
	fmt.Println(hasPathSum(ToTreeNode([]int{1, 2, 3}), 5) == false)
	fmt.Println(hasPathSum(ToTreeNode([]int{}), 0) == false)
}

// @lc code=end
