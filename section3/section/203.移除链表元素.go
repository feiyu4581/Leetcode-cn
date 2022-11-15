package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}

	prev, current := dummy, dummy.Next
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
			current = current.Next
		} else {
			prev = current
			current = current.Next
		}
	}

	return dummy.Next
}

func Test203() {
	// [1,2,3,4,5]
	fmt.Println(removeElements(ToListNode([]int{1, 2, 6, 3, 4, 5, 6}), 6).ToValues())
	// []
	fmt.Println(removeElements(ToListNode([]int{}), 1).ToValues())
	// []
	fmt.Println(removeElements(ToListNode([]int{7, 7, 7, 7}), 7).ToValues())
}

// @lc code=end
