package section

import "fmt"

/*
 * @lc app=leetcode.cn id=1463 lang=golang
 *
 * [1463] 摘樱桃 II
 */

// @lc code=start
func cherryPickup2(grid [][]int) int {
	row, column := len(grid), len(grid[0])

	dp := make([][][]int, row)
	for k := range dp {
		dp[k] = make([][]int, column+2)
		for i := range dp[k] {
			dp[k][i] = make([]int, column+2)
			for j := range dp[k][i] {
				dp[k][i][j] = -1
			}
		}
	}

	max := func(nums ...int) int {
		res := nums[0]
		for _, num := range nums {
			if num > res {
				res = num
			}
		}
		return res
	}

	dp[0][1][column] = grid[0][0] + grid[0][column-1]
	maxValue := 0
	for k := 1; k < row; k++ {
		for x1 := 1; x1 <= column; x1++ {
			for x2 := 1; x2 <= column; x2++ {
				dp[k][x1][x2] = max(
					dp[k-1][x1-1][x2-1],
					dp[k-1][x1-1][x2],
					dp[k-1][x1-1][x2+1],

					dp[k-1][x1][x2-1],
					dp[k-1][x1][x2],
					dp[k-1][x1][x2+1],

					dp[k-1][x1+1][x2-1],
					dp[k-1][x1+1][x2],
					dp[k-1][x1+1][x2+1],
				)

				if dp[k][x1][x2] >= 0 {
					if x1 == x2 {
						dp[k][x1][x2] += grid[k][x1-1]
					} else {
						dp[k][x1][x2] += grid[k][x1-1] + grid[k][x2-1]
					}

					if k == row-1 {
						maxValue = max(maxValue, dp[k][x1][x2])
					}
				}
			}
		}
	}

	return maxValue
}

// @lc code=end

func Test1463() {
	fmt.Println(cherryPickup2([][]int{
		{3, 1, 1},
		{2, 5, 1},
		{1, 5, 5},
		{2, 1, 1},
	}) == 24)

	fmt.Println(cherryPickup2([][]int{
		{1, 0, 0, 0, 0, 0, 1},
		{2, 0, 0, 0, 0, 3, 0},
		{2, 0, 9, 0, 0, 0, 0},
		{0, 3, 0, 5, 4, 0, 0},
		{1, 0, 2, 3, 0, 0, 6},
	}) == 28)

	fmt.Println(cherryPickup2([][]int{
		{1, 0, 0, 3},
		{0, 0, 0, 3},
		{0, 0, 3, 3},
		{9, 0, 3, 3},
	}) == 22)

	fmt.Println(cherryPickup2([][]int{
		{1, 1},
		{1, 1},
	}) == 4)

	fmt.Println(cherryPickup2([][]int{
		{0, 8, 7, 10, 9, 10, 0, 9, 6},
		{8, 7, 10, 8, 7, 4, 9, 6, 10},
		{8, 1, 1, 5, 1, 5, 5, 1, 2},
		{9, 4, 10, 8, 8, 1, 9, 5, 0},
		{4, 3, 6, 10, 9, 2, 4, 8, 10},
		{7, 3, 2, 8, 3, 3, 5, 9, 8},
		{1, 2, 6, 5, 6, 2, 0, 10, 0},
	}) == 96)
}
