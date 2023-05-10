package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=368 lang=golang
 *
 * [368] 最大整除子集
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	sort.Ints(nums)

	dp := make([]int, len(nums))
	parent := make([]int, len(nums))
	for i := range parent {
		parent[i] = i
		dp[i] = 0
	}

	maxIndex := 0
	for i := range nums {
		maxJ := -1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 {
				if maxJ == -1 || dp[j] > dp[maxJ] {
					maxJ = j
				}
			}

			if maxJ != -1 {
				dp[i] = dp[maxJ] + 1
				parent[i] = maxJ
				if dp[i] > dp[maxIndex] {
					maxIndex = i
				}
			}
		}
	}

	res := make([]int, dp[maxIndex]+1)
	current := maxIndex
	for i := dp[maxIndex]; i >= 0; i-- {
		res[i] = nums[current]
		current = parent[current]
	}
	return res
}

// @lc code=end

func Test368() {
	// [1, 2] or [1, 3]
	fmt.Println(largestDivisibleSubset([]int{1, 2, 3}))

	// [1, 2, 4, 8]
	fmt.Println(largestDivisibleSubset([]int{1, 2, 4, 8}))

	// [4,8,240]
	fmt.Println(largestDivisibleSubset([]int{4, 8, 10, 240}))
}
