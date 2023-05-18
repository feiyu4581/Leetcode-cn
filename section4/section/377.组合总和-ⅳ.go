package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=377 lang=golang
 *
 * [377] 组合总和 Ⅳ
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)

	cache := make(map[int]int, 0)
	var dfs func(int) int
	dfs = func(num int) int {
		if num == 0 {
			return 1
		}

		if _, ok := cache[num]; !ok {
			res := 0
			for i := 0; i < len(nums) && num >= nums[i]; i++ {
				res += dfs(num - nums[i])
			}

			cache[num] = res
		}

		return cache[num]
	}

	return dfs(target)
}

// @lc code=end

func Test377() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4) == 7)
	fmt.Println(combinationSum4([]int{9}, 3) == 0)
}
