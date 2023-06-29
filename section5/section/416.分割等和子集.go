package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
func canPartition(nums []int) bool {

	allNum := 0
	for _, num := range nums {
		allNum += num
	}

	if allNum%2 != 0 {
		return false
	}

	sort.Ints(nums)
	var dfs func(int, int) bool
	cache := make(map[[2]int]bool, 0)
	dfs = func(i int, current int) bool {
		if current == 0 {
			return true
		}

		if i < 0 {
			return false
		}

		key := [2]int{i, current}
		if _, ok := cache[key]; !ok {
			if nums[i] <= current && dfs(i-1, current-nums[i]) {
				cache[key] = true
				return true
			}

			cache[key] = dfs(i-1, current)
		}

		return cache[key]
	}

	return dfs(len(nums)-1, allNum/2)
}

// @lc code=end

func Test416() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}) == false)
}
