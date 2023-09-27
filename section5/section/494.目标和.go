package section

import (
	"fmt"
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	allSum := 0
	for _, num := range nums {
		allSum += num
	}

	if allSum < target {
		return 0
	}

	diffNum := allSum - target
	if diffNum%2 != 0 {
		return 0
	}

	matchNum := diffNum / 2

	sort.Ints(nums)

	cache := make(map[[2]int]int)
	var dfs func(index int, targetValue int) int
	dfs = func(index int, targetValue int) int {
		if targetValue == 0 {
			return 1
		}

		key := [2]int{index, targetValue}
		if _, ok := cache[key]; !ok {
			ans := 0
			for i := index; i < len(nums) && nums[i] <= targetValue; i++ {
				ans += dfs(i+1, targetValue-nums[i])
			}

			cache[key] = ans
		}

		return cache[key]
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			return dfs(i, matchNum) * int(math.Pow(2, float64(i)))
		}
	}

	return int(math.Pow(2, float64(len(nums))))
}

// @lc code=end

func Test494() {
	fmt.Println(findTargetSumWays([]int{0, 0}, 0) == 4)
	fmt.Println(findTargetSumWays([]int{1, 0}, 1) == 2)
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3) == 5)
	fmt.Println(findTargetSumWays([]int{1}, 1) == 1)
}
