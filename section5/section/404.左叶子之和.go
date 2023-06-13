package section

import "fmt"

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	var traverse func(*TreeNode, bool)

	val := 0
	traverse = func(node *TreeNode, isLeft bool) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			if isLeft {
				val += node.Val
			}
		} else {
			traverse(node.Left, true)
			traverse(node.Right, false)
		}
	}

	traverse(root, false)
	return val
}

// @lc code=end

func Test404() {
	fmt.Println(sumOfLeftLeaves(ToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})) == 24)
	fmt.Println(sumOfLeftLeaves(ToTreeNode([]int{1})) == 0)
}
