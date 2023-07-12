package section

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	choose := 0
	lastI := math.MinInt32
	for i := range intervals {
		if intervals[i][0] >= lastI {
			choose += 1
			lastI = intervals[i][1]
		}
	}

	return len(intervals) - choose
}

// @lc code=end

func Test435() {
	fmt.Println(eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}) == 1)

	fmt.Println(eraseOverlapIntervals([][]int{
		{1, 2},
		{1, 2},
		{1, 2},
	}) == 2)

	fmt.Println(eraseOverlapIntervals([][]int{
		{1, 2},
		{2, 3},
	}) == 0)
}
