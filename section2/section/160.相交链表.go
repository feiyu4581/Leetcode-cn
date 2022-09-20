package section

import "fmt"

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func computeNodeLength(head *ListNode) int {
	current := head
	length := 0
	for current != nil {
		length += 1
		current = current.Next
	}

	return length
}

func computeIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA.Next == nil && headB.Next == nil {
		if headA == headB {
			return headA
		}

		return nil
	}

	node := computeIntersectionNode(headA.Next, headB.Next)
	if node != nil && node == headA.Next && headA == headB {
		return headA
	}

	return node
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	leftLength, rightLength := computeNodeLength(headA), computeNodeLength(headB)
	if leftLength > rightLength {
		leftLength, rightLength = rightLength, leftLength
		headA, headB = headB, headA
	}

	diffLength := rightLength - leftLength
	for i := 0; i < diffLength; i++ {
		headB = headB.Next
	}

	return computeIntersectionNode(headA, headB)
}

func Test160() {
	fmt.Println(getIntersectionNode(ToListNode([]int{4, 1, 8, 4, 5}), ToListNode([]int{5, 6, 1, 8, 4, 5})).ToValues())
}

// @lc code=end
