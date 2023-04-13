package section

import "fmt"

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	dummyOdd, dummyEven := &ListNode{}, &ListNode{}

	odd, even := dummyOdd, dummyEven

	current := head
	isOdd := true
	for current != nil {
		if isOdd {
			odd.Next = current
			odd = odd.Next
		} else {
			even.Next = current
			even = even.Next
		}

		isOdd = !isOdd
		current = current.Next
	}

	odd.Next = dummyEven.Next
	even.Next = nil
	return dummyOdd.Next
}

// @lc code=end

func Test328() {
	// [1,3,5,2,4]
	fmt.Println(oddEvenList(ToListNode([]int{1, 2, 3, 4, 5})).ToValues())

	// [2,3,6,7,1,5,4]
	fmt.Println(oddEvenList(ToListNode([]int{2, 1, 3, 5, 6, 4, 7})).ToValues())
}
