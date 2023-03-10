package section

import "fmt"

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for n > 3 {
		allowN := 0
		for _, num := range []int{2, 3, 5} {
			if n%num == 0 {
				allowN = num
				break
			}
		}

		if allowN == 0 {
			return false
		}

		n = n / allowN
	}

	return true
}

// @lc code=end

func Test263() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(1))
	fmt.Println(isUgly(14) == false)
}
