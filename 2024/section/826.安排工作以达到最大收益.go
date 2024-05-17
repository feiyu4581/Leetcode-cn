package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=826 lang=golang
 *
 * [826] 安排工作以达到最大收益
 */

// @lc code=start
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	items := make([][2]int, len(difficulty))
	for i := range items {
		items[i] = [2]int{difficulty[i], profit[i]}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i][1] >= items[j][1]
	})

	res := 0
	for _, num := range worker {
		for i := 0; i < len(items); i++ {
			if num >= items[i][0] {
				res += items[i][1]
				break
			}
		}
	}

	return res
}

// @lc code=end

func Test826() {
	fmt.Println(maxProfitAssignment([]int{2, 4, 6, 8, 10}, []int{10, 20, 30, 40, 50}, []int{4, 5, 6, 7}) == 100)
	fmt.Println(maxProfitAssignment([]int{85, 47, 57}, []int{24, 66, 99}, []int{40, 25, 25}) == 0)
}
