package section

import "fmt"

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	// 收集所有的数据
	values := make([][]int, 0)
	// 当前行所有的节点，初始化的时候为第一行 root 节点数据
	currents := []*TreeNode{root}
	// 方向控制，默认正向遍历
	position := true
	for len(currents) > 0 {
		currentValues := make([]int, 0)

		var nextCurrents []*TreeNode
		// 遍历所有 root 节点
		for index, currentNode := range currents {
			// 如果是负向遍历的时候，将 index 改成从尾部开始
			if !position {
				index = len(currents) - index - 1
			}

			// 收集节点数据
			currentValues = append(currentValues, currents[index].Val)
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
		// 重置方向
		position = !position
	}

	return values
}

func Test103() {
	// [[3],[20,9],[15,7]]
	fmt.Println(zigzagLevelOrder(ToNode([]int{3, 9, 20, 0, 0, 15, 7})))

	// [[1]]
	fmt.Println(zigzagLevelOrder(ToNode([]int{1})))

	// []
	fmt.Println(zigzagLevelOrder(ToNode([]int{})))
}

// @lc code=end
