package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=274 lang=golang
 *
 * [274] H 指数
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	for i := range citations {
		if citations[i] >= len(citations)-i {
			return len(citations) - i
		}
	}

	return 0
}

// @lc code=end

func Test274() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}) == 3)
	fmt.Println(hIndex([]int{1, 3, 1}) == 1)
}
