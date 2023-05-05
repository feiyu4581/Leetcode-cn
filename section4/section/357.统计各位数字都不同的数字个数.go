package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 统计各位数字都不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	num := 1
	for i := 1; i <= n; i++ {
		tempNum := 9
		for j := 1; j < i; j++ {
			tempNum = tempNum * (10 - j)
		}

		num += tempNum
	}

	return num
}

// @lc code=end

func Test357() {
	fmt.Println(countNumbersWithUniqueDigits(2) == 91)
	fmt.Println(countNumbersWithUniqueDigits(0) == 1)
}
