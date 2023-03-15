package section

import "fmt"

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	current := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[current] = nums[i]
			current += 1
		}
	}

	for current < len(nums) {
		nums[current] = 0
		current += 1
	}
}

// @lc code=end

func Test283() {
	num := []int{0, 1, 0, 3, 12}
	moveZeroes(num)
	// [1,3,12,0,0]
	fmt.Println(num)
}
