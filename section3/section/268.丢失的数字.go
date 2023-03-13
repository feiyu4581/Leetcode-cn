package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	gotMax := false
	for i := range nums {
		num := nums[i]
		if num < 0 {
			num = -num - 1
		}

		if num == len(nums) {
			gotMax = true
			continue
		}

		nums[num] = -nums[num] - 1
	}

	if !gotMax {
		return len(nums)
	}

	for i := range nums {
		if nums[i] >= 0 {
			return i
		}
	}

	return 0
}

// @lc code=end

func Test268() {
	fmt.Println(missingNumber([]int{2, 0}) == 1)
	fmt.Println(missingNumber([]int{3, 0, 1}) == 2)
	fmt.Println(missingNumber([]int{0, 1}) == 2)
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}) == 8)
}
