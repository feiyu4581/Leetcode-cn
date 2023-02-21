package section

import "fmt"

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func Test226() {
	// [4,7,2,9,6,3,1]
	fmt.Println(invertTree(ToTreeNode([]int{4, 2, 7, 1, 3, 6, 9})).ToValues())
	// [2,3,1]
	fmt.Println(invertTree(ToTreeNode([]int{2, 1, 3})).ToValues())
	// []
	fmt.Println(invertTree(ToTreeNode([]int{})).ToValues())
}

// @lc code=end
