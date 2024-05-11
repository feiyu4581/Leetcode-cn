package section

import "fmt"

/*
 * @lc app=leetcode.cn id=2960 lang=golang
 *
 * [2960] 统计已测试设备
 */

// @lc code=start
func countTestedDevices(batteryPercentages []int) int {
	count := 0
	for _, battery := range batteryPercentages {
		if battery > count {
			count++
		}
	}
	return count
}

// @lc code=end

func Test2960() {
	fmt.Println(countTestedDevices([]int{1, 1, 2, 1, 3}) == 3)
	fmt.Println(countTestedDevices([]int{0, 1, 2}) == 2)
}
