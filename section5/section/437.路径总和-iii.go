package section

import "fmt"

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(*TreeNode, int, []int)
	res := 0
	dfs = func(node *TreeNode, sum int, pathValues []int) {
		if node == nil {
			return
		}

		currentSum := sum + node.Val
		if currentSum == targetSum {
			res += 1
		}

		for _, pathValue := range pathValues {
			currentSum -= pathValue
			if currentSum == targetSum {
				res += 1
			}
		}

		dfs(node.Left, sum+node.Val, append(pathValues, node.Val))
		dfs(node.Right, sum+node.Val, append(pathValues, node.Val))
	}

	dfs(root, 0, []int{})
	return res
}

// @lc code=end

func Test437() {
	fmt.Println(pathSum(ToTreeNode([]int{-2, 0, -3}), -3) == 1)
	fmt.Println(pathSum(ToTreeNode([]int{10, 5, -3, 3, 2, 0, 11, 3, -2, 0, 1}), 8) == 3)
	fmt.Println(pathSum(ToTreeNode([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 5, 1}), 22) == 3)
}
