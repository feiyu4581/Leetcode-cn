package section

import "fmt"

/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func countZeroAndOne(str string) (int, int) {
	m, n := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			m++
		} else {
			n++
		}
	}

	return m, n
}

func max474(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][][]int, len(strs)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = make([]int, n+1)
		}
	}

	for i := 1; i <= len(strs); i++ {
		zeroCount, oneCount := countZeroAndOne(strs[i-1])
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j >= zeroCount && k >= oneCount {
					dp[i][j][k] = max474(dp[i-1][j][k], dp[i-1][j-zeroCount][k-oneCount]+1)
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	return dp[len(strs)][m][n]
}

// @lc code=end

func Test474() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3) == 4)
	fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1) == 2)
}
