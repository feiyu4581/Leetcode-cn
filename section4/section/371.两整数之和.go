package section

import "fmt"

/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start
func getSum(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}

// @lc code=end

func Test371() {
	fmt.Println(getSum(1, 2) == 3)
	fmt.Println(getSum(2, 3) == 5)
}
