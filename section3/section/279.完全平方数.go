package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	if n <= 3 {
		return n
	}

	squares := make([]int, 0)
	squares = append(squares, 1)
	for i := 2; i <= n; i++ {
		if i*i < n {
			squares = append(squares, i*i)
		} else if i*i == n {
			return 1
		} else {
			break
		}
	}

	min := func(left, right int) int {
		if left < right {
			return left
		}

		return right
	}
	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 3
	for i := 4; i <= n; i++ {
		current := 0
		minValue := math.MaxInt32
		for current < len(squares) && squares[current] <= i {
			minValue = min(minValue, dp[i-squares[current]]+1)
			current += 1
		}

		dp[i] = minValue
	}

	return dp[n]
}

// @lc code=end

func Test279() {
	fmt.Println(numSquares(12) == 3)
	fmt.Println(numSquares(13) == 2)
}
