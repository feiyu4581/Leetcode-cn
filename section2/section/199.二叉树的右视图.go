package section

import "fmt"

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	roots := make([]*TreeNode, 0)
	depths := make([]int, 0)

	current := root
	depth := 1

	res := make([]int, 0)
	for len(roots) > 0 || current != nil {
		if current != nil {
			if depth > len(res) {
				res = append(res, current.Val)
			}

			roots = append(roots, current)
			depths = append(depths, depth)

			depth += 1
			current = current.Right
		} else {
			current = roots[len(roots)-1]
			roots = roots[:len(roots)-1]

			depth = depths[len(depths)-1]
			depths = depths[:len(depths)-1]

			current = current.Left
			depth += 1
		}
	}

	return res
}

func Test199() {
	// [1, 3, 4]
	fmt.Println(rightSideView(ToTreeNode([]int{1, 2, 3, 0, 5, 0, 4})))
	// [1, 3]
	fmt.Println(rightSideView(ToTreeNode([]int{1, 0, 3})))
}

// @lc code=end
