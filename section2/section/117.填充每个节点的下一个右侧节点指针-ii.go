package section

import "fmt"

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a NextNode.
 * type NextNode struct {
 *     Val int
 *     Left *NextNode
 *     Right *NextNode
 *     Next *NextNode
 * }
 */

func connect(root *NextNode) *NextNode {
	if root == nil {
		return root
	}

	currents := []*NextNode{root}
	for len(currents) > 0 {
		for index := range currents[:len(currents)-1] {
			currents[index].Next = currents[index+1]
		}

		nextCurrent := make([]*NextNode, 0, len(currents))
		for _, node := range currents {
			if node.Left != nil {
				nextCurrent = append(nextCurrent, node.Left)
			}

			if node.Right != nil {
				nextCurrent = append(nextCurrent, node.Right)
			}
		}
		currents = nextCurrent
	}

	return root
}

func Test117() {
	// [1,#,2,3,#,4,5,7,#]
	fmt.Println(connect(ToNode([]int{1, 2, 3, 4, 5, 0, 7})).ToValues())

	fmt.Println(connect(ToNode([]int{1, 2, 3, 4, 5, 0, 6, 7, 0, 0, 0, 0, 8})).ToValues())

	fmt.Println(connect(ToNode([]int{3, 9, 20, 0, 0, 15, 7})).ToValues())

	fmt.Println(connect(ToNode([]int{1, 12, 22, 13, 23, 33, 43, 14, 24, 34, 44, 54, 64, 0, 0, 15, 25})).ToValues())
}

// @lc code=end
