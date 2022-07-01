package section

import "fmt"

/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func Test108() {
	// [0,-3,9,-10,null,5] 或 0,-10,5,null,-3,null,9]
	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}).ToValues())

	// [1, 3]
	fmt.Println(sortedArrayToBST([]int{1, 3}).ToValues())
}

// @lc code=end
