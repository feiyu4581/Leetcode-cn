package section

import "fmt"

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for n >= 4 {
		if n%4 != 0 {
			return false
		}

		n = n / 4
	}

	return n == 1
}

// @lc code=end

func Test342() {
	fmt.Println(isPowerOfFour(16))
	fmt.Println(isPowerOfFour(5) == false)
	fmt.Println(isPowerOfFour(1))
}
