package section

import "fmt"

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a NextNode.
 * type NextNode struct {
 *     Val int
 *     Left *NextNode
 *     Right *NextNode
 *     Next *NextNode
 * }
 */

func connectNode1(left *NextNode, right *NextNode) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	connectNode1(left.Left, left.Right)
	connectNode1(right.Left, right.Right)
	connectNode1(left.Right, right.Left)
}

func connect1(root *NextNode) *NextNode {
	if root == nil {
		return root
	}
	connectNode1(root.Left, root.Right)
	return root
}

func Test116() {
	// 1,#,2,3,#,4,5,6,7,#
	fmt.Println(connect1(ToNode([]int{1, 2, 3, 4, 5, 6, 7})).ToValues())
}

// @lc code=end
