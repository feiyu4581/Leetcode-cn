package section

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	nodeMaps := make(map[*ListNode]struct{}, 0)
	current := head
	for current != nil {
		if _, ok := nodeMaps[current]; ok {
			return current
		}

		nodeMaps[current] = struct{}{}
		current = current.Next
	}

	return nil
}

// @lc code=end
