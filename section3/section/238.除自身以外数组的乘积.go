package section

import "fmt"

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	prefixs := make([]int, len(nums))
	suffixs := make([]int, len(nums))

	prefix, suffix := 1, 1
	for i := 0; i < len(nums); i++ {
		prefix = prefix * nums[i]
		prefixs[i] = prefix

		suffix = suffix * nums[len(nums)-i-1]
		suffixs[len(nums)-i-1] = suffix
	}

	nums[0], nums[len(nums)-1] = suffixs[1], prefixs[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		nums[i] = prefixs[i-1] * suffixs[i+1]
	}

	return nums
}

// @lc code=end

func Test238() {
	// [24,12,8,6]
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	// [0,0,9,0,0]
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
