package section

import "fmt"

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start
func findDuplicate(nums []int) int {
	abs := func(value int) int {
		if value < 0 {
			return -value
		}

		return value
	}

	defer func() {
		for i := range nums {
			nums[i] = abs(nums[i])
		}
	}()

	for i := range nums {
		num := abs(nums[i])
		if nums[num] < 0 {
			return num
		}

		nums[num] = -nums[num]
	}

	return nums[0]
}

// @lc code=end

func Test287() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}) == 2)
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}) == 3)
}
