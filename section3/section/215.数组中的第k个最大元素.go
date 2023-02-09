package section

import "fmt"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

func partition(nums []int) int {
	current, index := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= current {
			nums[i], nums[index] = nums[index], nums[i]
			index += 1
		}
	}

	if index == 0 {
		return index
	}

	nums[0], nums[index-1] = nums[index-1], nums[0]
	return index - 1
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	mid := partition(nums) + 1
	if mid == k {
		return nums[mid-1]
	} else if mid > k {
		return findKthLargest(nums[:mid-1], k)
	} else {
		return findKthLargest(nums[mid:], k-mid)
	}
}

func Test215() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) == 5)
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) == 4)
}

// @lc code=end
