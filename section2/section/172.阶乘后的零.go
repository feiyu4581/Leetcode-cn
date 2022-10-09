package section

import "fmt"

/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

func Test172() {
	//fmt.Println(trailingZeroes(3) == 0)
	//fmt.Println(trailingZeroes(5) == 1)
	//fmt.Println(trailingZeroes(0) == 0)
	fmt.Println(trailingZeroes(40) == 9)
}

// @lc code=end
