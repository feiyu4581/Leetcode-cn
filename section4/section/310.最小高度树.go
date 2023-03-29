package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=310 lang=golang
 *
 * [310] 最小高度树
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	nodes := make(map[int][]int, n)

	degrees := make(map[int]int, n)
	for i := 0; i < n; i++ {
		degrees[i] = 0
	}

	for _, edge := range edges {
		nodes[edge[0]] = append(nodes[edge[0]], edge[1])
		nodes[edge[1]] = append(nodes[edge[1]], edge[0])

		degrees[edge[1]] += 1
		degrees[edge[0]] += 1
	}

	current := make([]int, 0)
	for node, degree := range degrees {
		if degree == 1 {
			current = append(current, node)
		}
	}

	for len(degrees) > 2 {
		newCurrent := make([]int, 0)
		for _, node := range current {
			delete(degrees, node)
			for _, edge := range nodes[node] {
				if _, ok := degrees[edge]; ok {
					degrees[edge] -= 1
					if degrees[edge] == 1 {
						newCurrent = append(newCurrent, edge)
					}
				}
			}
		}

		current = newCurrent
	}

	res := make([]int, 0, len(degrees))
	for node := range degrees {
		res = append(res, node)
	}

	return res
}

// @lc code=end

func Test310() {
	// [3]
	fmt.Println(findMinHeightTrees(6, [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{3, 4},
		{4, 5},
	}))
	//  [1]
	fmt.Println(findMinHeightTrees(4, [][]int{
		{1, 0},
		{1, 2},
		{1, 3},
	}))

	// [3,4]
	fmt.Println(findMinHeightTrees(6, [][]int{
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 4},
		{5, 4},
	}))
}
