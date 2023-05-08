package section

import "fmt"

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start

var maxSqure367 = 1 << 16

func isPerfectSquare(num int) bool {
	if num <= 0 {
		return false
	}

	left, right := 0, num
	for left <= right {
		mid := left + (right-left)>>1
		if mid >= maxSqure367 {
			mid = maxSqure367 - 1
		}

		midNum := mid * mid
		if midNum == num {
			return true
		}

		if midNum > num {
			right = mid - 1
		} else {
			left = mid + 1

		}
	}

	return false
}

// @lc code=end

func Test367() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14) == false)
}
