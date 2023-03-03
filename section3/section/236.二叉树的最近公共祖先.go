package section

import "fmt"

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	return lowestCommonAncestor(root, p, q)
}

// @lc code=end

func Test236() {
	root := ToTreeNode([]int{3, 5, 1, 6, 2, -1, 8, 0, 0, 7, 4})
	fmt.Println(lowestCommonAncestor236(root, root.Left, root.Right).Val == 3)
}
