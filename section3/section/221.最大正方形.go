package section

import "fmt"

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for index := range dp {
		dp[index] = make([]int, len(matrix[0]))
	}

	min := func(num ...int) int {
		minN := num[0]
		for i := 1; i < len(num); i++ {
			if num[i] < minN {
				minN = num[i]
			}
		}

		return minN
	}

	maxN := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				if i > 0 && j > 0 {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				} else {
					dp[i][j] = 1
				}

				if dp[i][j] > maxN {
					maxN = dp[i][j]
				}
			}
		}
	}

	return maxN * maxN
}

func Test221() {
	fmt.Println(maximalSquare([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}) == 4)

	fmt.Println(maximalSquare([][]byte{
		{'0', '1'},
		{'1', '0'},
	}) == 1)

	fmt.Println(maximalSquare([][]byte{
		{'0'},
	}) == 0)
}

// @lc code=end
