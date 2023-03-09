package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	results := make([]string, 0)
	var dfs func(node *TreeNode, prefix []string)
	dfs = func(node *TreeNode, prefix []string) {
		prefix = append(prefix, fmt.Sprintf("%d", node.Val))
		if node.Left == nil && node.Right == nil {
			results = append(results, strings.Join(prefix, "->"))
			return
		}

		if node.Left != nil {
			dfs(node.Left, prefix[:])
		}

		if node.Right != nil {
			dfs(node.Right, prefix[:])
		}
	}

	dfs(root, make([]string, 0))
	return results
}

// @lc code=end

func Test257() {
	// ["1->2->5","1->3"]
	fmt.Println(binaryTreePaths(ToTreeNode([]int{1, 2, 3, 0, 5})))

	// ["1"]
	fmt.Println(binaryTreePaths(ToTreeNode([]int{1})))
}
