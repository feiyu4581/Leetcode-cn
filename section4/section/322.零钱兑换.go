package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}

	sort.Ints(coins)
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			if i-coins[j] < 0 {
				continue
			}
			if i-coins[j] == 0 {
				dp[i] = 1
				break
			}

			if diff := dp[i-coins[j]]; diff >= 1 {
				if dp[i] == -1 || diff+1 < dp[i] {
					dp[i] = diff + 1
				}
			}
		}
	}

	return dp[amount]
}

// @lc code=end

func Test322() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11) == 3)
	fmt.Println(coinChange([]int{2}, 3) == -1)
	fmt.Println(coinChange([]int{1}, 0) == 0)
}
