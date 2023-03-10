package section

import "fmt"

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	min := func(left, right int) int {
		if left < right {
			return left
		}

		return right
	}

	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		if dp[i] == dp[p2]*2 {
			p2++
		}
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
	}

	return dp[n]
}

// @lc code=end

func Test264() {
	fmt.Println(nthUglyNumber(10) == 12)
	fmt.Println(nthUglyNumber(1) == 1)
}
