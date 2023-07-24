package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start

func abs442(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		if nums[abs442(num)-1] < 0 {
			res = append(res, abs442(num))
		} else {
			nums[abs442(num)-1] = -nums[abs442(num)-1]
		}
	}
	return res
}

// @lc code=end

func Test442() {
	// [2, 3]
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))

	// [1]
	fmt.Println(findDuplicates([]int{1, 1, 2}))

	// []
	fmt.Println(findDuplicates([]int{1}))
}
