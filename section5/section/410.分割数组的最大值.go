package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] 分割数组的最大值
 */

// @lc code=start

func max410(left, right int) int {
	if left > right {
		return left
	}

	return right
}

func splitArray(nums []int, k int) int {
	m := len(nums)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	num := 0
	for i := 0; i < m; i++ {
		num += nums[i]
		dp[i][1] = num
	}

	for step := 2; step <= k; step++ {
		for i := step - 1; i < m; i++ {
			minValue := math.MaxInt32
			num = 0
			for j := i; j >= step-1; j-- {
				num += nums[j]
				if diff := max410(num, dp[j-1][step-1]); diff < minValue {
					minValue = diff
				}
			}
			dp[i][step] = minValue
		}

	}

	return dp[m-1][k]
}

// @lc code=end

func Test410() {
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2) == 9)
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2) == 18)
	fmt.Println(splitArray([]int{1, 4, 4}, 3) == 4)
}
