package section

import "fmt"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	currentCount, currentValue := 0, nums[0]
	start, end := 0, 1
	for end <= len(nums) {
		if currentValue >= target {
			if end-start < currentCount || currentCount == 0 {
				currentCount = end - start
			}

			if currentValue-target >= nums[start] && start < end {
				currentValue -= nums[start]
				start += 1
			} else if end < len(nums) {
				currentValue += nums[end]
				end += 1
			} else {
				break
			}
		} else if end < len(nums) {
			currentValue += nums[end]
			end += 1
		} else {
			break
		}
	}

	return currentCount
}

func Test209() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}) == 2)
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}) == 1)
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}) == 0)
}

// @lc code=end
