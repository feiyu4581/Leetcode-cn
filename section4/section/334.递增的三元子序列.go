package section

import "fmt"

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	dpMax := make([]int, len(nums))
	maxIndex := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 || nums[i] >= nums[maxIndex] {
			maxIndex = i
		}

		dpMax[i] = maxIndex
	}

	minIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if i == 0 || nums[i] <= nums[minIndex] {
			minIndex = i
		}

		if i > minIndex && i < dpMax[i] {
			return true
		}
	}

	return false
}

// @lc code=end

func Test334() {
	fmt.Println(increasingTriplet([]int{1, 2, 2, 1}) == false)
	fmt.Println(increasingTriplet([]int{1, 1, 1, 1, 1}) == false)
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}) == false)
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
