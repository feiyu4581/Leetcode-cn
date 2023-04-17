package section

import "fmt"

/*
 * @lc app=leetcode.cn id=330 lang=golang
 *
 * [330] 按要求补齐数组
 */

// @lc code=start
func minPatches(nums []int, n int) int {
	patch := 0
	current := 1

	i := 0
	for current <= n {
		if i < len(nums) && nums[i] <= current {
			current += nums[i]
			i += 1
		} else {
			current = current << 1
			patch += 1
		}
	}

	return patch
}

// @lc code=end

func Test330() {
	fmt.Println(minPatches([]int{1, 3}, 6) == 1)
	fmt.Println(minPatches([]int{1, 5, 10}, 20) == 2)
	fmt.Println(minPatches([]int{1, 2, 2}, 5) == 0)
}
