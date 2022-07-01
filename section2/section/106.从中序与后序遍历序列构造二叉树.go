package section

import "fmt"

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	headValue := postorder[len(postorder)-1]
	splitIndex := 0
	for index, inorderValue := range inorder {
		if headValue == inorderValue {
			splitIndex = index
			break
		}
	}

	return &TreeNode{
		Val:   headValue,
		Left:  buildTree(inorder[0:splitIndex], postorder[0:splitIndex]),
		Right: buildTree(inorder[splitIndex+1:], postorder[splitIndex:len(postorder)-1]),
	}
}

func Test106() {
	// [3,9,20,null,null,15,7]
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}).ToValues())

	// [-1]
	fmt.Println(buildTree([]int{-1}, []int{-1}).ToValues())
}

// @lc code=end
