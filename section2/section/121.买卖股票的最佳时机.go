package section

import "fmt"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	start := prices[0]
	for _, price := range prices {
		if price > start {
			if profit < (price - start) {
				profit = price - start
			}
		} else {
			start = price
		}

	}

	return profit
}

func Test121() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println(maxProfit2([]int{7, 6, 4, 3, 1}) == 0)
}

// @lc code=end
