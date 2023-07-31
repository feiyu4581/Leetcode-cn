package section

import "fmt"

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}

	ans := make([]int, 0)
	for index, num := range nums {
		if num > 0 {
			ans = append(ans, index+1)
		}
	}
	return ans
}

func Test448() {
	// [5, 6]
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// [2]
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

// @lc code=end
