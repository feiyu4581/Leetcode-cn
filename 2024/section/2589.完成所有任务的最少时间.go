package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2589 lang=golang
 *
 * [2589] 完成所有任务的最少时间
 */

// @lc code=start
func findMinimumTime(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool { return tasks[i][1] < tasks[j][1] })
	vis := [2010]int{}
	ans := 0
	for _, task := range tasks {
		start, end, duration := task[0], task[1], task[2]
		for _, x := range vis[start : end+1] {
			duration -= x
		}
		for i := end; i >= start && duration > 0; i-- {
			if vis[i] == 0 {
				vis[i] = 1
				duration--
				ans++
			}
		}
	}
	return ans
}

// @lc code=end

func Test2589() {
	fmt.Println(findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}}) == 2)
	fmt.Println(findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}}) == 4)
	fmt.Println(findMinimumTime([][]int{{14, 20, 5}, {2, 18, 7}, {6, 14, 1}, {3, 16, 3}}) == 7)
}
