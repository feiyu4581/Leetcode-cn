package section

import "fmt"

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
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

	visited := make(map[*TreeNode]struct{}, 0)
	for len(stacks) > 0 {
		node := pop()

		if node.Right == nil && node.Left == nil {
			nums = append(nums, node.Val)
			continue
		}

		if _, ok := visited[node]; ok {
			nums = append(nums, node.Val)
			continue
		}
		stacks = append(stacks, node)

		if node.Right != nil {
			stacks = append(stacks, node.Right)
		}

		if node.Left != nil {
			stacks = append(stacks, node.Left)
		}

		visited[node] = struct{}{}
	}

	return nums
}

func Test145() {
	// [3, 2, 1]
	fmt.Println(postorderTraversal(ToTreeNode([]int{1, 0, 2, 3})))

	// []
	fmt.Println(postorderTraversal(ToTreeNode([]int{})))

	// [1]
	fmt.Println(postorderTraversal(ToTreeNode([]int{1})))
}

// @lc code=end
