package section

import "fmt"

/*
 * @lc app=leetcode.cn id=495 lang=golang
 *
 * [495] 提莫攻击
 */

// @lc code=start
func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0

	for i := range timeSeries {
		if i+1 < len(timeSeries) && timeSeries[i+1]-timeSeries[i] < duration {
			ans += timeSeries[i+1] - timeSeries[i]
		} else {
			ans += duration
		}
	}

	return ans
}

// @lc code=end

func Test495() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2) == 4)
	fmt.Println(findPoisonedDuration([]int{1, 2}, 2) == 3)
}
