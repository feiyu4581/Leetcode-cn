package section

import "fmt"

/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start

var positions = [][2]int{
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
}

func longestIncreasingPath(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if dp[i][j] > 0 {
			return dp[i][j]
		}

		maxNum := 0
		for _, position := range positions {
			newI, newJ := i+position[0], j+position[1]
			if newI >= 0 && newJ >= 0 && newI < len(matrix) && newJ < len(matrix[0]) && matrix[i][j] < matrix[newI][newJ] {
				if num := dfs(newI, newJ); num > maxNum {
					maxNum = num
				}
			}
		}

		dp[i][j] = maxNum + 1
		return dp[i][j]
	}

	maxCount := 0
	for i := range matrix {
		for j := range matrix[i] {
			if num := dfs(i, j); num > maxCount {
				maxCount = num
			}
		}
	}

	return maxCount
}

// @lc code=end

func Test329() {
	fmt.Println(longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}) == 4)

	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}) == 4)

	fmt.Println(longestIncreasingPath([][]int{
		{1},
	}) == 1)
}
