package section

import "fmt"

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	headValue := preorder[0]
	splitIndex := 0
	for index, inorderValue := range inorder {
		if headValue == inorderValue {
			splitIndex = index
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:splitIndex+1], inorder[0:splitIndex]),
		Right: buildTree(preorder[splitIndex+1:], inorder[splitIndex+1:]),
	}
}

func Test105() {
	// [3,9,20,null,null,15,7]
	fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).ToValues())

	// [-1]
	fmt.Println(buildTree([]int{-1}, []int{-1}).ToValues())
}

// @lc code=end
