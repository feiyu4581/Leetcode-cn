package section

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ToListNode(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head

	for _, value := range values[1:] {
		current.Next = &ListNode{
			Val: value,
		}
		current = current.Next
	}
	return head
}

func (node *ListNode) String() string {
	return fmt.Sprintf("[List]%d", node.Val)
}

func (node *ListNode) ToValues() []int {
	values := make([]int, 0)
	current := node

	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	return values
}
