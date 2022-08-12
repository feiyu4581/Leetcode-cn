package section

import "fmt"

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stacks := []*TreeNode{root}

	nums := make([]int, 0)

	pop := func() *TreeNode {
		node := stacks[len(stacks)-1]
		stacks = stacks[:len(stacks)-1]
		return node
	}

	for len(stacks) > 0 {
		node := pop()
		nums = append(nums, node.Val)
		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}

		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}

	}

	return nums
}

func Test144() {
	// [1, 2, 3]
	fmt.Println(preorderTraversal(ToTreeNode([]int{1, 0, 2, 3})))

	// []
	fmt.Println(preorderTraversal(ToTreeNode([]int{})))

	// [1]
	fmt.Println(preorderTraversal(ToTreeNode([]int{1})))
}

// @lc code=end
