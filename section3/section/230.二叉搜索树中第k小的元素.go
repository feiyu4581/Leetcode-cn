package section

import "fmt"

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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

func dfsKthSmallest(root *TreeNode, k int) (int, bool) {
	if root == nil {
		return 0, false
	}

	if root.Left == nil && root.Right == nil {
		if k == 1 {
			return root.Val, true
		}
		return 1, false
	}

	if leftCount, ok := dfsKthSmallest(root.Left, k); ok {
		return leftCount, ok
	} else if leftCount+1 == k {
		return root.Val, true
	} else if rightCount, ok := dfsKthSmallest(root.Right, k-leftCount-1); ok {
		return rightCount, ok
	} else {
		return leftCount + rightCount + 1, false
	}
}

func kthSmallest(root *TreeNode, k int) int {
	if num, ok := dfsKthSmallest(root, k); ok {
		return num
	}

	return 0
}

// @lc code=end

func Test230() {
	fmt.Println(kthSmallest(ToTreeNode([]int{3, 1, 4, 0, 2}), 1) == 1)
	fmt.Println(kthSmallest(ToTreeNode([]int{5, 3, 6, 2, 4, 0, 0, 1}), 3) == 3)
}
