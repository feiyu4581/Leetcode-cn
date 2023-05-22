package section

import (
	"fmt"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	root   *ListNode
	Length int
}

func Constructor382(head *ListNode) Solution {
	length := 0
	current := head
	for current != nil {
		length += 1
		current = current.Next
	}

	return Solution{
		root:   head,
		Length: length,
	}
}

func (this *Solution) GetRandom() int {
	index := rand.Intn(this.Length)
	current := this.root
	for i := 1; i <= index; i++ {
		current = current.Next
	}

	return current.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

func Test382() {
	obj := Constructor382(ToListNode([]int{1, 2, 3}))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.GetRandom())
}
