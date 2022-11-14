package section

import "fmt"

/*
 * @lc app=leetcode.cn id=201 lang=golang
 *
 * [201] 数字范围按位与
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	shift := 0
	for left < right {
		left, right = left>>1, right>>1
		shift += 1
	}

	return left << shift

}

func Test201() {
	fmt.Println(rangeBitwiseAnd(5, 7) == 4)
	fmt.Println(rangeBitwiseAnd(0, 0) == 0)
	fmt.Println(rangeBitwiseAnd(1, 2147483647) == 0)
}

// @lc code=end
