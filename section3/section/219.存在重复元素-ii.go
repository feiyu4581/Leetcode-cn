package section

import "fmt"

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	counter := make(map[int]int, len(nums))
	for index := range nums {
		if lastIndex, ok := counter[nums[index]]; ok && index-lastIndex <= k {
			return true
		}

		counter[nums[index]] = index
	}

	return false
}

func Test219() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2) == false)
}

// @lc code=end
