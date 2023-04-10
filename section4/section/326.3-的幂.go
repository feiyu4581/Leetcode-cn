package section

import "fmt"

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3 的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	if n <= 1 {
		return n == 1
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}

		n = n / 3
	}

	return true
}

// @lc code=end

func Test326() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0) == false)
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45) == false)
}
