package section

import "fmt"

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = k % len(nums)

	tempNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		tempNums[(i+k)%len(nums)] = nums[i]
	}

	for index := range nums {
		nums[index] = tempNums[index]
	}
}

func Test189() {
	// [5,6,7,1,2,3,4]
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)

	// [3,99,-1,-100]
	nums = []int{-1, -100, 3, 99}
	rotate(nums, 2)
	fmt.Println(nums)
}

// @lc code=end
