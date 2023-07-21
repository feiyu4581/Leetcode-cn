package section

import "fmt"

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

// @lc code=start
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}

	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if mid*(mid+1)/2 == n {
			return mid
		} else if mid*(mid+1)/2 < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left - 1
}

// @lc code=end

func Test441() {
	fmt.Println(arrangeCoins(5) == 2)
	fmt.Println(arrangeCoins(8) == 3)
}
