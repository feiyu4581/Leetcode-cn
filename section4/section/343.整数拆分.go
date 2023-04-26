package section

import "fmt"

/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	max := func(i, j int) int {
		if i > j {
			return i
		}

		return j
	}
	for i := 2; i <= n; i++ {
		maxValue := 0
		for j := 1; j < i; j++ {
			if num := max(j, dp[j]) * max(i-j, dp[i-j]); num > maxValue {
				maxValue = num
			}
		}

		dp[i] = maxValue
	}

	return dp[n]
}

// @lc code=end

func Test343() {
	fmt.Println(integerBreak(2) == 1)
	fmt.Println(integerBreak(10) == 36)
}
