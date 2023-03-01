package section

import "fmt"

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func recursiveIsPalindrome(last *ListNode, first *ListNode) (bool, *ListNode) {
	if last.Next == nil {
		return last.Val == first.Val, first.Next
	}

	if ok, nextNode := recursiveIsPalindrome(last.Next, first); ok {
		return last.Val == nextNode.Val, nextNode.Next
	}

	return false, nil
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	ok, _ := recursiveIsPalindrome(head, head)
	return ok
}

// @lc code=end

func Test234() {
	fmt.Println(isPalindrome(ToListNode([]int{1, 2, 2, 1})) == true)
	fmt.Println(isPalindrome(ToListNode([]int{1, 2})) == false)
	fmt.Println(isPalindrome(ToListNode([]int{1})) == true)
}
