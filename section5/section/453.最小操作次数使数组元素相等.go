package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	ans := 0
	sort.Ints(nums)

	lastH := nums[0]
	for i := 1; i < len(nums); i++ {
		currentH := ans + nums[i]
		ans = ans + currentH - lastH
		lastH = currentH
	}

	return ans
}

// @lc code=end

func Test453() {
	fmt.Println(minMoves([]int{1, 2, 3, 4}))
	fmt.Println(minMoves([]int{1, 1, 1, 1, 1}))
	fmt.Println(minMoves([]int{1, 2, 3}) == 3)
	fmt.Println(minMoves([]int{1, 1, 1}) == 0)
}

//f[n] = d[n]
//f[n + 1] = d[n] + x[n + 1] + d[n] - (x[n] + d[n - 1])
