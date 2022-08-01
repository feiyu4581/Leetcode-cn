package section

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node2 struct {
	Val    int
	Next   *Node2
	Random *Node2
}

func copyRandomListByCache(head *Node2, cache map[*Node2]*Node2) *Node2 {
	if head == nil {
		return nil
	}

	if node, ok := cache[head]; ok {
		return node
	}

	node := &Node2{Val: head.Val}
	cache[head] = node
	node.Next = copyRandomListByCache(head.Next, cache)
	node.Random = copyRandomListByCache(head.Random, cache)

	return node
}

func copyRandomList(head *Node2) *Node2 {
	return copyRandomListByCache(head, make(map[*Node2]*Node2, 0))
}

func Test138() {

}

// @lc code=end
