package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=473 lang=golang
 *
 * [473] 火柴拼正方形
 */

// @lc code=start

func chooseSequareSide(matchsticiks []int, target int, start int) ([]int, bool) {
	res := make([]int, 0)
	for i := start; i < len(matchsticiks); i++ {
		if matchsticiks[i] > target {
			continue
		}

		if matchsticiks[i] == target {
			return append(res, i), true
		}

		if matchsticiks[i] < target {
			if nextSides, ok := chooseSequareSide(matchsticiks, target-matchsticiks[i], i+1); ok {
				return append(append(res, i), nextSides...), true
			}
		}
	}

	return res, target == 0
}

func makesquare(matchsticks []int) bool {
	if len(matchsticks) < 4 {
		return false
	}

	allSum := 0
	for _, stick := range matchsticks {
		allSum += stick
	}

	if allSum%4 != 0 {
		return false
	}

	side := allSum / 4
	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] >= matchsticks[j]
	})

	if matchsticks[0] > side {
		return false
	}

	edges := [4]int{0, 0, 0, 0}
	var dfs func(int) bool
	dfs = func(i int) bool {
		if i == len(matchsticks) {
			return true
		}

		for j := range edges {
			if matchsticks[i]+edges[j] <= side {
				edges[j] += matchsticks[i]
				if dfs(i + 1) {
					return true
				}
				edges[j] -= matchsticks[i]
			}
		}

		return false
	}

	return dfs(0)
}

// @lc code=end

func Test473() {
	fmt.Println(makesquare([]int{13, 11, 1, 8, 6, 7, 8, 8, 6, 7, 8, 9, 8}))
	fmt.Println(makesquare([]int{2, 2, 2, 2, 2, 6}) == false)
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}) == false)
}
