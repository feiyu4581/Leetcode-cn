package section

import (
	"fmt"
	"math"
)

type Flag int

const (
	SUCCESS Flag = iota
	Failed
	OK
)

/*
 * @lc app=leetcode.cn id=174 lang=golang
 *
 * [174] 地下城游戏
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	min := func(num1, num2 int) int {
		if num1 > num2 {
			return num2
		}

		return num1
	}

	max := func(num1, num2 int) int {
		if num1 > num2 {
			return num1
		}

		return num2
	}

	dp[m-1][n], dp[m][n-1] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}

	return dp[0][0]
}

func Test174() {
	fmt.Println(calculateMinimumHP([][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}) == 7)

	fmt.Println(calculateMinimumHP([][]int{
		{1, -3, 3},
		{0, -2, 0},
		{-3, -3, -3},
	}) == 3)
}

// @lc code=end
