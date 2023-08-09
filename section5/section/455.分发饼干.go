package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	current := 0

	for i := 0; i < len(s) && current < len(g); i++ {
		if s[i] >= g[current] {
			ans++
			current++
		}
	}

	return ans
}

// @lc code=end

func Test455() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}) == 1)
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}) == 2)
}
