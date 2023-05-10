package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=363 lang=golang
 *
 * [363] 矩形区域不超过 K 的最大数值和
 */

// @lc code=start

type Node363 struct {
	val   int
	left  *Node363
	right *Node363
}

func (node *Node363) AddNode(val int) {
	current := node
	for {
		if current.val == val {
			return
		}

		if current.val > val {
			if current.left == nil {
				current.left = &Node363{val: val}
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = &Node363{val: val}
				return
			}

			current = current.right
		}
	}
}

func (node *Node363) FindMaxNode(val int) *Node363 {
	current := node
	var nextNode *Node363

	for current != nil {
		if current.val == val {
			nextNode = current
			break
		}

		if current.val > val {
			nextNode = current
			current = current.left
		} else {
			current = current.right
		}
	}

	return nextNode
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	res := math.MinInt32
	for start := range matrix {
		columns := make([]int, len(matrix[start]))
		for end := start; end < len(matrix); end++ {
			s := 0
			tree := Node363{val: 0}
			for i, num := range matrix[end] {
				columns[i] += num

				s += columns[i]

				nextNode := tree.FindMaxNode(s - k)
				if nextNode != nil && s-nextNode.val > res {
					res = s - nextNode.val
				}

				tree.AddNode(s)
			}
		}
	}

	return res
}

// @lc code=end

func Test363() {

	fmt.Println(maxSumSubmatrix([][]int{
		{5, -4, -3, 4},
		{-3, -4, 4, 5},
		{5, 1, 5, -4},
	}, 3) == 2)

	fmt.Println(maxSumSubmatrix([][]int{
		{1, 0, 1},
		{0, -2, 3},
	}, 2) == 2)

	fmt.Println(maxSumSubmatrix([][]int{
		{2, 2, -1},
	}, 0) == -1)

	fmt.Println(maxSumSubmatrix([][]int{
		{2, 2, -1},
	}, 3) == 3)
}
