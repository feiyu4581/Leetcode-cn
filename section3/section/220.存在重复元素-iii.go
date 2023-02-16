package section

import "fmt"

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	abs := func(left int) int {
		if left > 0 {
			return left
		}

		return -left
	}
	for i := range nums {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if abs(nums[j]-nums[i]) <= valueDiff {
				return true
			}
		}
	}

	return false
}

func Test220() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3) == false)
}

// @lc code=end
