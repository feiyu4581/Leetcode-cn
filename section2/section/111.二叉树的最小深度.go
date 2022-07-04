package section

import "fmt"

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if left == 0 || right == 0 {
		return left + right + 1
	}

	if left <= right {
		return left + 1
	}

	return right + 1
}

func Test111() {
	fmt.Println(minDepth(ToNode([]int{3, 9, 20, 0, 0, 15, 7})) == 2)
	fmt.Println(minDepth(ToNode([]int{2, 0, 3, 0, 4, 0, 5, 0, 6})) == 5)
}

// @lc code=end
