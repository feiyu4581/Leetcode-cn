package section

import "fmt"

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	splitDigits := func(digit int) int {
		current := 0

		for digit >= 10 {
			current += digit % 10
			digit = digit / 10
		}

		return current + digit
	}

	for num >= 10 {
		num = splitDigits(num)
	}

	return num
}

// @lc code=end

func Test258() {
	fmt.Println(addDigits(38) == 2)
	fmt.Println(addDigits(0) == 0)
}
