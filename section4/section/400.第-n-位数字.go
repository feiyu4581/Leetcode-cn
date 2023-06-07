package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=400 lang=golang
 *
 * [400] 第 N 位数字
 */

// @lc code=start
func findNthDigit(n int) int {
	if n <= 9 {
		return n
	}
	n -= 9

	minValue, maxValue, step := 10, 100, 2
	for {
		maxNums := (maxValue - minValue) * step
		if n <= maxNums {
			numBytes := []byte(strconv.Itoa((n-1)/step + minValue))
			return int(numBytes[(n-1)%step] - '0')
		} else {
			n -= maxNums
			minValue = maxValue
			maxValue = maxValue * 10
			step += 1
		}
	}
}

// @lc code=end

func Test400() {
	fmt.Println(findNthDigit(3) == 3)
	fmt.Println(findNthDigit(11) == 0)
	fmt.Println(findNthDigit(30) == 2)
	fmt.Println(findNthDigit(1000) == 3)
}
