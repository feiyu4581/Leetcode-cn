package section

import "fmt"

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fastNode, slowNode := head.Next, head
	for fastNode != slowNode {
		if fastNode == nil || fastNode.Next == nil {
			return false
		}

		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	return true
}

func Test141() {
	fmt.Println(hasCycle(ToListNode([]int{3, 2, 0, -4})))
	fmt.Println(hasCycle(ToListNode([]int{1, 2})))
	fmt.Println(hasCycle(ToListNode([]int{1})) == false)
}

// @lc code=end
