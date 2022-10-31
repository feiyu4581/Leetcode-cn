package section

import "fmt"

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	max := func(values ...int) int {
		value := 0
		for index := range values {
			if values[index] > value {
				value = values[index]
			}
		}

		return value
	}

	if len(nums) <= 2 {
		return max(nums...)
	}

	first, second := nums[0], nums[1]
	third := max(nums[2]+first, second)

	res := third
	for index := 3; index < len(nums); index++ {
		res = max(nums[index]+first, max(nums[index]+second), third)
		first = second
		second = third
		third = res
	}

	return res
}

func Test198() {
	fmt.Println(rob([]int{1, 2, 3, 1}) == 4)
	fmt.Println(rob([]int{2, 7, 9, 3, 1}) == 12)
	fmt.Println(rob([]int{2, 1, 1, 2}) == 4)
	fmt.Println(rob([]int{1, 1, 1}) == 2)
}

// @lc code=end
