package section

import "fmt"

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] <= nums[right] {
			return nums[left]
		}

		mid := left + ((right - left) >> 1)
		if mid == left {
			return nums[right]
		}

		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return nums[0]
}

func Test153() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}) == 1)
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}) == 0)
	fmt.Println(findMin([]int{11, 13, 15, 17}) == 11)
}

// @lc code=end
