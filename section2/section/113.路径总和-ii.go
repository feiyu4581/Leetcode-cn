package section

import "fmt"

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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

func computePathSum(root *TreeNode, targetSum int, current []int) (results [][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			output := make([]int, 0, len(current)+1)
			for _, num := range current {
				output = append(output, num)
			}

			output = append(output, root.Val)
			results = append(results, output)
			return
		}

		return
	}

	current = append(current, root.Val)
	results = append(results, computePathSum(root.Left, targetSum-root.Val, current)...)
	results = append(results, computePathSum(root.Right, targetSum-root.Val, current)...)
	return
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	return computePathSum(root, targetSum, []int{})
}

func Test113() {
	// [[5,4,11,2],[5,8,4,5]]
	fmt.Println(pathSum(ToTreeNode([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}), 22))

	// []
	fmt.Println(pathSum(ToTreeNode([]int{1, 2, 3}), 5))

	// []
	fmt.Println(pathSum(ToTreeNode([]int{1, 2}), 0))
}

// @lc code=end
