package section

import "fmt"

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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

func collectSumNums(root *TreeNode, current int) []int {
	if root.Left == nil && root.Right == nil {
		return []int{current*10 + root.Val}
	}

	results := make([]int, 0)
	if root.Left != nil {
		results = append(results, collectSumNums(root.Left, current*10+root.Val)...)
	}

	if root.Right != nil {
		results = append(results, collectSumNums(root.Right, current*10+root.Val)...)
	}

	return results
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	allNums := collectSumNums(root, 0)
	result := 0
	for _, num := range allNums {
		result += num
	}

	return result
}

func Test129() {
	fmt.Println(sumNumbers(ToTreeNode([]int{1, 2, 3})) == 25)
	fmt.Println(sumNumbers(ToTreeNode([]int{4, 9, 0, 5, 1})) == 1026)
}

// @lc code=end
