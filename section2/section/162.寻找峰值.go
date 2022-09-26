package section

import "fmt"

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func Test162() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}) == 2)
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) == 5) // 或者 1
}

// @lc code=end
