package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	maxNum := math.MinInt32
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	secondNum := math.MinInt32
	findSecondNum := false
	for _, num := range nums {
		if num < maxNum && num >= secondNum {
			secondNum = num
			findSecondNum = true
		}
	}

	if !findSecondNum {
		return maxNum
	}

	thirdNum := math.MinInt32
	findThirdNum := false
	for _, num := range nums {
		if num < secondNum && num >= thirdNum {
			thirdNum = num
			findThirdNum = true
		}
	}

	if findThirdNum {
		return thirdNum
	}

	return maxNum
}

// @lc code=end

func Test414() {
	fmt.Println(thirdMax([]int{3, 2, 1}) == 1)
	fmt.Println(thirdMax([]int{2, 2, 3, 1}) == 1)
	fmt.Println(thirdMax([]int{1, 2}) == 2)
}
