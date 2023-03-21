package section

import "fmt"

/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	maxLis := 1
	for i := 0; i < len(nums); i++ {
		lis := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > lis {
				lis = dp[j] + 1
			}
		}
		dp[i] = lis
		if dp[i] > maxLis {
			maxLis = dp[i]
		}
	}

	return maxLis
}

// @lc code=end

func Test300() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}) == 4)
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}) == 4)
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}) == 1)
}
