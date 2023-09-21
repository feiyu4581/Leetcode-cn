package section

import "fmt"

/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// @lc code=start
func predictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for step := 0; step < len(nums); step++ {
		for start := 0; start+step < len(nums); start++ {
			if step == 0 {
				dp[start][start] = nums[start]
			} else {
				dp[start][start+step] = max(nums[start]-dp[start+1][start+step], nums[start+step]-dp[start][start+step-1])
			}
		}
	}

	return dp[0][len(nums)-1] >= 0
}

// @lc code=end

func Test486() {
	fmt.Println(predictTheWinner([]int{1, 5, 2}) == false)
	fmt.Println(predictTheWinner([]int{1, 5, 233, 7}) == true)
}
