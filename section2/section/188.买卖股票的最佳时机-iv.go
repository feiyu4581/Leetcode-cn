package section

import "fmt"

/*
 * @lc app=leetcode.cn id=188 lang=golang
 *
 * [188] 买卖股票的最佳时机 IV
 */

// @lc code=start
func maxProfit(k int, prices []int) int {
	dp := make([][2]int, k)
	for i := 0; i < k; i++ {
		dp[i][0] = -prices[0]
	}

	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}

		return num2
	}

	for _, price := range prices {
		for i := 0; i < k; i++ {
			if i == 0 {
				dp[i][0] = max(dp[i][0], -price)
				dp[i][1] = max(dp[i][1], price+dp[i][0])
			} else {
				dp[i][0] = max(dp[i][0], dp[i-1][1]-price)
				dp[i][1] = max(dp[i][1], price+dp[i][0])
			}
		}
	}

	return dp[k-1][1]
}

func Test188() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}) == 2)
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}) == 7)
}

// @lc code=end
