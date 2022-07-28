package section

import "fmt"

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber2(nums []int) int {
	result := 0
	for _, num := range nums {
		result = result ^ num
	}

	return result
}

func Test136() {
	fmt.Println(singleNumber2([]int{2, 2, 1}) == 1)
	fmt.Println(singleNumber2([]int{4, 1, 2, 1, 2}) == 4)
}

// @lc code=end
