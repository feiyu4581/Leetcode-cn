package section

import "fmt"

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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

func replaceNode(node *TreeNode, source *TreeNode, target *TreeNode) {
	if node == nil {
		return
	}
	if node.Left == source {
		node.Left = target
	}
	if node.Right == source {
		node.Right = target
	}
}

func deleteMinNode(node *TreeNode) *TreeNode {
	var parentNode *TreeNode
	subNode := node
	for subNode.Left != nil {
		parentNode = subNode
		subNode = subNode.Left
	}

	replaceNode(parentNode, subNode, subNode.Right)
	return subNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	var parent *TreeNode
	current := root

	for current != nil {
		if current.Val == key {
			break
		}
		parent = current
		if current.Val > key {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	var nextParentNode *TreeNode
	if current != nil {
		// 1. 没有叶子节点，直接删除
		if current.Left == nil && current.Right == nil {
			replaceNode(parent, current, nil)
		} else if current.Left != nil && current.Right != nil {
			subNode := deleteMinNode(current.Right)
			replaceNode(parent, current, subNode)
			subNode.Left = current.Left
			if subNode != current.Right {
				subNode.Right = current.Right
			}

			nextParentNode = subNode
		} else {
			if current.Left != nil {
				replaceNode(parent, current, current.Left)
				nextParentNode = current.Left
			} else {
				replaceNode(parent, current, current.Right)
				nextParentNode = current.Right
			}
		}
	}

	if parent == nil {
		return nextParentNode
	}
	return root
}

// @lc code=end

func Test450() {
	// [5,4,6,2,7]
	fmt.Println(deleteNode(ToTreeNode([]int{5, 3, 6, 2, 4, 0, 7}), 3).ToValues())
	// [5,3,6,2,4,7]
	fmt.Println(deleteNode(ToTreeNode([]int{5, 3, 6, 2, 4, 0, 7}), 0).ToValues())
	// []
	fmt.Println(deleteNode(ToTreeNode([]int{}), 0).ToValues())
	// [6,3,7,2,4]
	fmt.Println(deleteNode(ToTreeNode([]int{5, 3, 6, 2, 4, 0, 7}), 5).ToValues())
}
