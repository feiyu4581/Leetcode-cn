package section

import "fmt"

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev, current := head, head.Next
	prev.Next = nil
	for current != nil {
		nextCurrent := current.Next
		current.Next = prev

		prev = current
		current = nextCurrent
	}

	return prev
}

func Test206() {
	// [5,4,3,2,1]
	fmt.Println(reverseList(ToListNode([]int{1, 2, 3, 4, 5})).ToValues())

	// [2, 1]
	fmt.Println(reverseList(ToListNode([]int{1, 2})).ToValues())

	// []
	fmt.Println(reverseList(ToListNode([]int{})).ToValues())
}

// @lc code=end
