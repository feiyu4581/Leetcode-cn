package section

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=462 lang=golang
 *
 * [462] 最小操作次数使数组元素相等 II
 */

// @lc code=start
func minMoves2(nums []int) int {
	sort.Ints(nums)
	average := nums[len(nums)/2]

	ans := 0
	for _, num := range nums {
		ans += int(math.Abs(float64(num - average)))
	}

	return ans
}

// @lc code=end

func Test462() {
	fmt.Println(minMoves2([]int{1, 2, 3}) == 2)
	fmt.Println(minMoves2([]int{1, 10, 2, 9}) == 16)
}
