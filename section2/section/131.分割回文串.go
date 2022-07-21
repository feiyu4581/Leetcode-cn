package section

import "fmt"

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

func backtrackPartition(s string, index int, currents []string, dp [][]bool) [][]string {
	if index == len(s) {
		partitions := make([]string, 0, len(currents))
		partitions = append(partitions, currents...)
		return [][]string{partitions}
	}

	results := make([][]string, 0)
	for i := index; i < len(s); i++ {
		if dp[index][i] {
			currents = append(currents, s[index:i+1])
			results = append(results, backtrackPartition(s, i+1, currents, dp)...)
			currents = currents[:len(currents)-1]
		}
	}

	return results
}

func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for index := range dp {
		dp[index] = make([]bool, len(s))
		dp[index][index] = true
		if index < len(dp)-1 && s[index] == s[index+1] {
			dp[index][index+1] = true
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}

	return backtrackPartition(s, 0, []string{}, dp)
}

func Test131() {
	// [["a","a","b"],["aa","b"]]
	fmt.Println(partition("aab"))
	// [["a"]]
	fmt.Println(partition("a"))
}

// @lc code=end
