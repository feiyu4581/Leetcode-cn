package section

import "fmt"

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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

func rangeTreeNode(root *TreeNode, index int, result [][]int) {
	if root != nil {
		rangeTreeNode(root.Left, index-1, result)
		result[index] = append(result[index], root.Val)
		rangeTreeNode(root.Right, index-1, result)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	depth := maxDepth(root)
	result := make([][]int, depth)
	rangeTreeNode(root, depth-1, result)
	return result
}

func Test107() {
	// [[15,7],[9,20],[3]]
	fmt.Println(levelOrderBottom(ToTreeNode([]int{3, 9, 20, 0, 0, 15, 7})))

	// [[1]]
	fmt.Println(levelOrderBottom(ToTreeNode([]int{1})))

	// []
	fmt.Println(levelOrderBottom(ToTreeNode([]int{})))
}

// @lc code=end
