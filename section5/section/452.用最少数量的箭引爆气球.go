package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			ans++
			end = points[i][1]
		} else {
			if points[i][1] < end {
				end = points[i][1]
			}
		}
	}

	return ans
}

// @lc code=end

func Test452() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}) == 2)
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}) == 4)
	fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}) == 2)
	fmt.Println(findMinArrowShots([][]int{
		{3, 9},
		{7, 12},
		{3, 8},
		{6, 8},
		{9, 10},
		{2, 9},
		{0, 9},
		{3, 9},
		{0, 6},
		{2, 8},
	}) == 2)
}
