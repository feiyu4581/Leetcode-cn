package section

import "fmt"

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}

		if total%3 != 0 {
			ans |= 1 << i
		}
	}

	return int(ans)
}

func Test137() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}) == 3)
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}) == 99)
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}) == -4)
}

// @lc code=end
