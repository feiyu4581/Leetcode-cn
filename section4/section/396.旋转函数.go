package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=396 lang=golang
 *
 * [396] 旋转函数
 */

// @lc code=start
func maxRotateFunction(nums []int) int {
	current := 0
	allSum := 0
	for i := 0; i < len(nums); i++ {
		current += i * nums[i]
		allSum += nums[i]
	}

	res := current
	for i := 1; i < len(nums); i++ {
		nextZero := nums[len(nums)-i]
		current = current + allSum - len(nums)*nextZero
		if current > res {
			res = current
		}
	}

	return res
}

// @lc code=end

func Test396() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}) == 26)
	fmt.Println(maxRotateFunction([]int{100}) == 0)
}
