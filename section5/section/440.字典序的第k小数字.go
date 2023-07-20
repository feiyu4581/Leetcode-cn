package section

/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 */

// @lc code=start

func min440(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getSteps(n, current int) int {
	res := 0
	first, last := current, current+1
	for first <= n {
		res = res + (min440(n+1, last) - first)
		first *= 10
		last *= 10
	}

	return res
}

func findKthNumber(n int, k int) int {
	res := 1

	k -= 1
	for k > 0 {
		steps := getSteps(n, res)
		if steps <= k {
			k -= steps
			res += 1
		} else {
			k -= 1
			res = res * 10
		}
	}

	return res
}

// @lc code=end

func Test440() {
	n := 128
	//findKthNumber(n, 6)
	for i := 1; i <= n; i++ {
		findKthNumber(n, i)
	}

}
