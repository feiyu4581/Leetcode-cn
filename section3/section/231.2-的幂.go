package section

import "fmt"

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	for i := 0; i < 32; i++ {
		if n == 1 {
			return true
		}

		if n&1 == 1 {
			return false
		}

		n = n >> 1
	}

	return false
}

func Test231() {
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(3) == false)
	fmt.Println(isPowerOfTwo(4))
	fmt.Println(isPowerOfTwo(5) == false)
}

// @lc code=end
