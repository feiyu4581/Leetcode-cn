package section

import "fmt"

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func findFastMid(head, tail *ListNode) *ListNode {
	fastNode, slowNode := head.Next, head.Next
	// 使用快慢指针找到中间值
	for fastNode != tail && fastNode.Next != tail {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	if slowNode.Next == tail {
		return head
	}

	mid := slowNode.Next
	slowNode.Next = mid.Next
	mid.Next = head

	return mid
}

func fastSortList(head, tail *ListNode) *ListNode {
	if head == nil || head == tail || head.Next == tail {
		return head
	}

	// 寻找中间值
	head = findFastMid(head, tail)
	mid := head
	start := head
	for start.Next != tail {
		// 所有小于中间值的节点移动到最左边
		if start.Next.Val < mid.Val {
			current := start.Next
			start.Next = current.Next
			current.Next = head
			head = current
		} else {
			start = start.Next
		}
	}

	head = fastSortList(head, mid)
	if mid.Next != nil {
		mid.Next = fastSortList(mid.Next, tail)
	}

	return head
}

func sortList(head *ListNode) *ListNode {
	// 快速排序
	return fastSortList(head, nil)
}

func Test148() {
	nums := make([]int, 0, 50000)
	for i := 0; i < 50000; i++ {
		nums = append(nums, i+1)
	}
	sortList(ToListNode(nums))
	fmt.Println(sortList(ToListNode([]int{4, 2, 1, 3})).ToValues())
	fmt.Println(sortList(ToListNode([]int{-1, 5, 3, 4, 0})).ToValues())
}

// @lc code=end
