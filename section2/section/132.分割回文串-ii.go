package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=132 lang=golang
 *
 * [132] 分割回文串 II
 */

// @lc code=start

func backtrackMinCut(s string, index int, cache map[int]int, dp [][]bool) int {
	if index == len(s) {
		return 0
	}

	if dp[index][len(s)-1] {
		return 0
	}

	if num, ok := cache[index]; ok {
		return num
	}

	minNum := math.MaxInt
	for i := index; i < len(s); i++ {
		if dp[index][i] {
			num := backtrackMinCut(s, i+1, cache, dp) + 1
			if minNum > num {
				minNum = num
			}
		}
	}

	cache[index] = minNum
	return minNum
}

func minCut(s string) int {
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

	return backtrackMinCut(s, 0, make(map[int]int, 0), dp)
}

func Test132() {
	fmt.Println(minCut("aab") == 1)
	fmt.Println(minCut("a") == 0)
	fmt.Println(minCut("ab") == 1)
}

// @lc code=end
