package section

import "fmt"

/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)

	var dfs func(index int, prefix []int)
	dfs = func(index int, prefix []int) {
		if len(prefix) >= 2 {
			newNums := append(make([]int, 0, len(prefix)), prefix...)
			res = append(res, newNums)
		}

		if index >= len(nums) {
			return
		}

		visits := make(map[int]struct{})
		for i := index; i < len(nums); i++ {
			if _, ok := visits[nums[i]]; !ok {
				visits[nums[i]] = struct{}{}
				if len(prefix) == 0 || prefix[len(prefix)-1] <= nums[i] {
					dfs(i+1, append(prefix, nums[i]))
				}
			}
		}
	}

	dfs(0, make([]int, 0))
	return res
}

// @lc code=end

func Test491() {
	// [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
	fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	// [[4, 4]]
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
}
