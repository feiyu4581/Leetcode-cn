package section

func guess(num int) int {
	return 0
}

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)>>1
		switch guess(mid) {
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		case 0:
			return mid
		}
	}

	return -1
}

// @lc code=end
