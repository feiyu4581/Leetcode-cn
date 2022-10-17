package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start

func maxProfit4(prices []int) int {
	buy1, buy2, sell1, sell2 := -prices[0], -prices[0], 0, 0
	max := func(left, right int) int {
		if left > right {
			return left
		}

		return right
	}
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}

	return sell2
}

func Test123() {
	fmt.Println(maxProfit4([]int{3, 5, 3, 4, 1, 4, 5, 0, 7, 8, 5, 6, 9, 4, 1}) == 13)
	fmt.Println(maxProfit4([]int{8, 6, 4, 3, 3, 2, 3, 5, 8, 3, 8, 2, 6}) == 11)
	fmt.Println(maxProfit4([]int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}) == 13)
	fmt.Println(maxProfit4([]int{3, 3, 5, 0, 0, 3, 1, 4}) == 6)
	fmt.Println(maxProfit4([]int{1, 2, 3, 4, 5}) == 4)
	fmt.Println(maxProfit4([]int{7, 6, 4, 3, 1}) == 0)
	fmt.Println(maxProfit4([]int{1}) == 0)
}

// @lc code=end
