package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=479 lang=golang
 *
 * [479] 最大回文数乘积
 */

// @lc code=start
func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}

	upper := int(math.Pow10(n)) - 1
	for left := upper; ; left-- {
		p := left
		for x := left; x > 0; x /= 10 {
			p = p*10 + (x % 10)
		}

		for ans := upper; ans*ans >= p; ans-- {
			if p%ans == 0 {
				return p % 1337
			}
		}
	}
}

// @lc code=end

func Test479() {
	fmt.Println(largestPalindrome(2) == 987)
	fmt.Println(largestPalindrome(1) == 9)
}
