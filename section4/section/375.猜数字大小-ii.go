package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=375 lang=golang
 *
 * [375] 猜数字大小 II
 */

// @lc code=start
func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}

	for step := 0; step < n; step++ {
		for i := 1; i <= n-step; i++ {
			j := i + step
			if step == 0 {
				dp[i][j] = 0
			} else if step == 1 {
				dp[i][j] = i
			} else {
				minValue := math.MaxInt32
				for k := i; k <= j; k++ {
					value := k
					leftValue, rightValue := 0, 0
					if i <= k-1 {
						leftValue = dp[i][k-1]
					}

					if k+1 <= j {
						rightValue = dp[k+1][j]
					}

					if leftValue > rightValue {
						value += leftValue
					} else {
						value += rightValue
					}

					if value < minValue {
						minValue = value
					}
				}
				dp[i][j] = minValue
			}
		}
	}

	return dp[1][n]
}

// @lc code=end

func Test375() {
	fmt.Println(getMoneyAmount(10) == 16)
	fmt.Println(getMoneyAmount(1) == 0)
	fmt.Println(getMoneyAmount(2) == 1)
}
