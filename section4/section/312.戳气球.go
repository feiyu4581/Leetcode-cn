package section

import "fmt"

/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */

// @lc code=start

func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	n := len(nums) - 1
	for i := 2; i <= n; i++ {
		for left := 0; left <= n-i; left++ {
			right := left + i
			maxVal := 0
			for mid := left + 1; mid < right; mid++ {
				if sumVal := dp[left][mid] + dp[mid][right] + nums[left]*nums[mid]*nums[right]; sumVal > maxVal {
					maxVal = sumVal
				}
			}

			dp[left][right] = maxVal
		}
	}

	return dp[0][n]
}

// @lc code=end

func Test312() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}) == 167)
	fmt.Println(maxCoins([]int{1, 5}) == 10)
}
