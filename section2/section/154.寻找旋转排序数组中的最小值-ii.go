package section

import "fmt"

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			return nums[left]
		}

		if nums[left] < nums[right] {
			return nums[left]
		}

		mid := left + ((right - left) >> 1)
		if mid == left {
			return nums[right]
		}

		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[left] == nums[right] {
			right -= 1
			continue
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[0]
}

func Test154() {
	fmt.Println(findMin([]int{3, 3, 1}) == 1)
	fmt.Println(findMin([]int{1, 3, 5}) == 1)
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}) == 0)
}

// @lc code=end
