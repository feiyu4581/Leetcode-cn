package section

import "fmt"

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func recursionSortedListToBST(head *ListNode, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}

	slowNode, fastNode := head, head
	for fastNode != tail && fastNode.Next != tail {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	return &TreeNode{
		Val:   slowNode.Val,
		Left:  recursionSortedListToBST(head, slowNode),
		Right: recursionSortedListToBST(slowNode.Next, tail),
	}
}
func sortedListToBST(head *ListNode) *TreeNode {
	return recursionSortedListToBST(head, nil)
}

func Test109() {
	// 0,-3,9,-10,null,5
	fmt.Println(sortedListToBST(ToListNode([]int{-10, -3, 0, 5, 9})).ToValues())
}

// @lc code=end
