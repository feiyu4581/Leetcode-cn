package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit3(prices []int) int {
	total := 0
	if len(prices) == 0 {
		return total
	}

	current := math.MaxInt
	for index := range prices {
		if prices[index] > current {
			if index+1 >= len(prices) {
				total += prices[index] - current
			} else if prices[index+1] < prices[index] {
				total += prices[index] - current
				current = prices[index]
			}
		} else {
			current = prices[index]
		}
	}

	return total
}

func Test122() {
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4}) == 7)
	fmt.Println(maxProfit3([]int{1, 2, 3, 4, 5}) == 4)
	fmt.Println(maxProfit3([]int{7, 6, 4, 3, 1}) == 0)
}

// @lc code=end
