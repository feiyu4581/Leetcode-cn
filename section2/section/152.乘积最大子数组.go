package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := i; j < len(nums); j++ {
			product = product * nums[j]
			if product > max {
				max = product
			}
		}
	}

	return max
}

func Test152() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}) == 6)
	fmt.Println(maxProduct([]int{-2, 0, -1}) == 0)
}

// @lc code=end
