package section

import "fmt"

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1C := make(chan *ListNode, 0)
	l2C := make(chan *ListNode, 0)

	var dfs func(node *ListNode, c chan *ListNode, index int)
	dfs = func(node *ListNode, c chan *ListNode, index int) {
		if node == nil {
			return
		}

		dfs(node.Next, c, index+1)
		c <- node

		if index == 0 {
			close(c)
		}
	}

	go dfs(l1, l1C, 0)
	go dfs(l2, l2C, 0)

	getValue := func(c chan *ListNode) (int, bool) {
		node, ok := <-c
		if ok {
			return node.Val, true
		}

		return 0, false
	}

	var current *ListNode
	carry := 0
	for {
		left, leftOk := getValue(l1C)
		right, rightOk := getValue(l2C)

		if !leftOk && !rightOk {
			break
		}

		value := left + right + carry
		if value >= 10 {
			value -= 10
			carry = 1
		} else {
			carry = 0
		}

		current = &ListNode{
			Val:  value,
			Next: current,
		}
	}

	if carry > 0 {
		current = &ListNode{
			Val:  carry,
			Next: current,
		}
	}

	return current
}

// @lc code=end

func Test445() {
	// [7,8,0,7]
	fmt.Println(addTwoNumbers(ToListNode([]int{7, 2, 4, 3}), ToListNode([]int{5, 6, 4})).ToValues())
	// [1,0,0,0,0,0,0,0,0,0,0]
	fmt.Println(addTwoNumbers(ToListNode([]int{9, 9, 9, 9, 9, 9, 9}), ToListNode([]int{1})).ToValues())
	// [0]
	fmt.Println(addTwoNumbers(ToListNode([]int{0}), ToListNode([]int{0})).ToValues())
	// [1, 0]
	fmt.Println(addTwoNumbers(ToListNode([]int{5}), ToListNode([]int{5})).ToValues())
}
