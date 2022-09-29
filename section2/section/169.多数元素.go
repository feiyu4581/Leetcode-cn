package section

import "fmt"

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := make(map[int]int, 0)
	maxValue, maxLength := 0, 0
	for _, num := range nums {
		count[num] += 1
		if count[num] > maxLength {
			maxLength = count[num]
			maxValue = num
		}
	}

	return maxValue
}

func Test169() {
	fmt.Println(majorityElement([]int{3, 2, 3}) == 3)
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) == 2)
}

// @lc code=end
