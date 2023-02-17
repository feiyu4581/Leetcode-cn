package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode, int, int) bool

	maxDepth, lastPosition, getEnd := 1, 0, 0
	dfs = func(node *TreeNode, depth int, position int) bool {
		maxDepth = depth
		if node.Left == nil {
			getEnd = depth
			return false
		}

		if node.Left != nil && node.Right == nil {
			lastPosition = 2*position - 1
			return true
		}

		if getEnd == depth {
			lastPosition = 2 * position
			return true
		}

		if dfs(node.Right, depth+1, 2*position) {
			return true
		}

		if dfs(node.Left, depth+1, 2*position-1) {
			return true
		}

		return false
	}

	dfs(root, 1, 1)
	return int(math.Pow(2, float64(maxDepth))) - 1 + lastPosition
}

func Test222() {
	fmt.Println(countNodes(ToTreeNode([]int{1, 2, 3, 4, 5, 6, 7})) == 7)
	fmt.Println(countNodes(ToTreeNode([]int{1, 2, 3, 4, 5})) == 5)
	fmt.Println(countNodes(ToTreeNode([]int{1, 2, 3})) == 3)
	fmt.Println(countNodes(ToTreeNode([]int{1, 2, 3, 4, 5, 6})) == 6)
	fmt.Println(countNodes(ToTreeNode([]int{})) == 0)
	fmt.Println(countNodes(ToTreeNode([]int{1})) == 1)
}

// @lc code=end
