package section

import "fmt"

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isEqual(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && isEqual(left.Left, right.Right) && isEqual(left.Right, right.Left)
}

//使用迭代方式来进行镜像校验
func isSymmetricByRange(root *TreeNode) bool {
	currents := []*TreeNode{root}
	for len(currents) > 0 {
		for i := 0; i < (len(currents) / 2); i++ {
			compare := currents[len(currents)-i-1]
			if currents[i] != nil && compare != nil {
				if currents[i].Val != compare.Val {
					return false
				}
			} else {
				if currents[i] != compare {
					return false
				}
			}
		}

		nextCurrents := make([]*TreeNode, 0, len(currents)*2)
		for _, current := range currents {
			if current != nil {
				nextCurrents = append(nextCurrents, current.Left, current.Right)
			}
		}

		currents = nextCurrents
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isEqual(root.Left, root.Right)
}

func Test101() {
	fmt.Println(isSymmetric(ToNode([]int{1, 2, 2, 3, 4, 4, 3})))
	fmt.Println(!isSymmetric(ToNode([]int{1, 2, 2, 0, 3, 0, 3})))
}

// @lc code=end
