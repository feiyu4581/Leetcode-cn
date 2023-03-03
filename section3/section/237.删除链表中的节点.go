package section

import "fmt"

/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Next, node.Val = node.Next.Next, node.Next.Val
}

// @lc code=end
func Test237() {
	node := ToListNode([]int{4, 5, 1, 9})
	deleteNode(node.Next)
	fmt.Println(node.ToValues())
}
