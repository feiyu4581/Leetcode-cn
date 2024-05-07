package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=741 lang=golang
 *
 * [741] 摘樱桃
 */

// @lc code=start
func cherryPickup(grid [][]int) int {
	length := len(grid)
	dp := make([][][]int, 2*length-1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]int, length)
			for k := 0; k < length; k++ {
				dp[i][j][k] = math.MinInt
			}
		}
	}

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	fetch := func(i, x1, x2 int) int {
		if i < 0 || x1 < 0 || x2 < 0 {
			return math.MinInt
		}

		return dp[i][x1][x2]
	}

	dp[0][0][0] = grid[0][0]
	for i := 1; i < 2*length-1; i++ {
		for x1 := 0; x1 < length && x1 <= i; x1++ {
			y1 := i - x1
			if y1 >= length {
				continue
			}

			for x2 := 0; x2 < length && x2 <= i; x2++ {
				y2 := i - x2
				if y2 >= length {
					continue
				}

				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[i][x1][x2] = math.MinInt
				} else {
					dp[i][x1][x2] = max(
						max(fetch(i-1, x1-1, x2-1), fetch(i-1, x1-1, x2)),
						max(fetch(i-1, x1, x2-1), fetch(i-1, x1, x2)),
					)

					if x1 == x2 {
						dp[i][x1][x2] += grid[x1][y1]
					} else {
						dp[i][x1][x2] += grid[x1][y1] + grid[x2][y2]
					}
				}
			}
		}
	}

	return max(dp[2*length-2][length-1][length-1], 0)
}

// @lc code=end

func Test741() {
	fmt.Println(cherryPickup([][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}) == 5)  // 5
	fmt.Println(cherryPickup([][]int{{1, 1, -1}, {1, -1, 1}, {-1, 1, 1}}) == 0) // 0
	fmt.Println(cherryPickup([][]int{{1, -1, 1}, {-1, 1, 1}, {1, 1, 1}}) == 0)
	fmt.Println(cherryPickup([][]int{
		{1, -1, 1, -1, 1, 1, 1, 1, 1, -1},
		{-1, 1, 1, -1, -1, 1, 1, 1, 1, 1},
		{1, 1, 1, -1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{-1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, -1, 1, 1, 1, 1, -1, 1, 1, 1},
		{1, 1, 1, -1, 1, 1, -1, 1, 1, 1},
		{1, -1, 1, -1, -1, 1, 1, 1, 1, 1},
		{1, 1, -1, -1, 1, 1, 1, -1, 1, -1},
		{1, 1, -1, 1, 1, 1, 1, 1, 1, 1},
	}) == 0)
}
