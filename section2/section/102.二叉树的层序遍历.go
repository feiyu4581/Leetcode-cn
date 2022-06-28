package section

import "fmt"

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 收集所有的数据
	values := make([][]int, 0)
	// 当前行所有的节点，初始化的时候为第一行 root 节点数据
	currents := []*TreeNode{root}
	for len(currents) > 0 {
		currentValues := make([]int, 0)

		var nextCurrents []*TreeNode
		// 遍历所有 root 节点
		for _, currentNode := range currents {
			// 收集节点数据
			currentValues = append(currentValues, currentNode.Val)
			// 将子节点放到下一次遍历数据中
			if currentNode.Left != nil {
				nextCurrents = append(nextCurrents, currentNode.Left)
			}

			if currentNode.Right != nil {
				nextCurrents = append(nextCurrents, currentNode.Right)
			}
		}

		values = append(values, currentValues)
		currents = nextCurrents
	}

	return values
}

func Test102() {
	// [[3],[9,20],[15,7]]
	fmt.Println(levelOrder(ToNode([]int{3, 9, 20, 0, 0, 15, 7})))

	// [[1]]
	fmt.Println(levelOrder(ToNode([]int{1})))

	// []
	fmt.Println(levelOrder(ToNode([]int{})))
}

// @lc code=end
