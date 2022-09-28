package section

import "fmt"

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	left, right := 1, len(numbers)
	for left < right {
		total := numbers[left-1] + numbers[right-1]
		if total == target {
			return []int{left, right}
		} else if total > target {
			right -= 1
		} else {
			left += 1
		}
	}

	return []int{}
}

func Test167() {
	// [1, 2]
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	// [1, 3]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	// [1, 2]
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

// @lc code=end
