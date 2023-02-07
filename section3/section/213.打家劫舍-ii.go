package section

import "fmt"

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func robLast(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[0][0] = 0
	dp[0][1] = 0

	max := func(left, right int) int {
		if left > right {
			return left
		}

		return right
	}

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return dp[len(nums)-1][1]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([][2]int, len(nums))
	dp[0][0] = 0
	dp[0][1] = nums[0]

	max := func(left, right int) int {
		if left > right {
			return left
		}

		return right
	}

	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[len(nums)-1][0], robLast(nums))
}

func Test213() {
	fmt.Println(rob([]int{2, 3, 2}) == 3)
	fmt.Println(rob([]int{1, 2, 3, 1}) == 4)
	fmt.Println(rob([]int{1, 2, 3}) == 3)
	fmt.Println(rob([]int{1}) == 1)
}

// @lc code=end
