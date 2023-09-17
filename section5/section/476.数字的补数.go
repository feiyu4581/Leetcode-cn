package section

import "fmt"

/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 */

// @lc code=start
func findComplement(num int) int {
	count := 0
	ans := 0
	for num != 0 {
		if num&1 == 0 {
			ans += 1 << count
		}

		count += 1
		num = num >> 1
	}

	return ans
}

// @lc code=end

func Test476() {
	fmt.Println(findComplement(5) == 2)
	fmt.Println(findComplement(1) == 0)
}
