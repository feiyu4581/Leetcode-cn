package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=354 lang=golang
 *
 * [354] 俄罗斯套娃信封问题
 */

// @lc code=start
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}

		return envelopes[i][0] < envelopes[j][0]
	})

	var f []int
	for _, e := range envelopes {
		h := e[1]
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

// @lc code=end

func Test354() {
	fmt.Println(maxEnvelopes([][]int{
		{2, 1},
		{4, 1},
		{6, 2},
		{8, 3},
		{10, 5},
		{12, 8},
		{14, 13},
		{16, 21},
		{18, 34},
		{20, 55},
	}) == 9)
	fmt.Println(maxEnvelopes([][]int{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}) == 3)
	fmt.Println(maxEnvelopes([][]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}) == 1)
}
