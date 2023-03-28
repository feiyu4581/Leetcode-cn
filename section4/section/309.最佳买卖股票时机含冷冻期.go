package section

import "fmt"

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	dp := make([][2]int, len(prices))

	max := func(left, right int) int {
		if left < right {
			return right
		}

		return left
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	dp[1][1] = max(dp[0][1], -prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
	}

	return dp[len(prices)-1][0]
}

// @lc code=end

func Test309() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}) == 3)
	fmt.Println(maxProfit([]int{1}) == 0)
}
