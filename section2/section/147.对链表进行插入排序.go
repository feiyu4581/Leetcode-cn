package section

import "fmt"

/*
 * @lc app=leetcode.cn id=147 lang=golang
 *
 * [147] 对链表进行插入排序
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	sortTail := head
	current := head.Next

	// [head,...,sorTail] 为已排序数据
	// current 为待插入数据
	for current != nil {
		nextCurrent := current.Next
		sortTail.Next = nextCurrent

		// 如果待插入数据小于头数据，那么将其插入到头节点，并且更新头结点
		if head.Val > current.Val {
			current.Next = head
			head = current
		} else {
			start := head
			// 找到满足 start.Val <= current.Val < start.Next.Val 的start节点，以 sortTail 为结束区间
			for start != sortTail && start.Next.Val <= current.Val {
				start = start.Next
			}

			// 如果已经找到了 sortTail 还没有满足的节点，那么 current 认为是有序的，将其插入到 sortTail 后，并且更新 sortTail
			if start == sortTail {
				start.Next = current
				sortTail = current
			} else {
				// 否则将其插入到中间节点，此时不需要更新 sortTail
				current.Next = start.Next
				start.Next = current
			}
		}

		current = nextCurrent
	}

	return head
}

func Test147() {
	fmt.Println(insertionSortList(ToListNode([]int{4, 2, 1, 3})).ToValues())
	fmt.Println(insertionSortList(ToListNode([]int{-1, 5, 3, 4, 0})).ToValues())
}

// @lc code=end
