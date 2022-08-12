package section

import "fmt"

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	nodeList := make([]*ListNode, 0)
	current := head
	for current != nil {
		nodeList = append(nodeList, current)
		current = current.Next
	}

	left, right := 1, len(nodeList)-1

	current = head
	for left <= right {
		current.Next = nodeList[right]
		current = current.Next

		if left < right {
			current.Next = nodeList[left]
			current = current.Next
		}

		current.Next = nil
		left += 1
		right -= 1
	}
}

func Test143() {
	node := ToListNode([]int{1, 2, 3, 4})
	reorderList(node)
	fmt.Println(node.ToValues())

	node = ToListNode([]int{1, 2, 3, 4, 5})
	reorderList(node)
	fmt.Println(node.ToValues())
}

// @lc code=end
