package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=233 lang=golang
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
func countDigitOne(n int) int {
	var recursive func(int) int
	recursive = func(num int) int {
		if num == 0 {
			return 0
		}

		if num <= 9 {
			return 1
		}

		current := 1
		for i := 1; i < 100; i++ {
			limit := int(math.Pow10(i))
			if num >= limit*10 {
				current = limit + 10*current
			} else {
				lastNum := num / limit
				if lastNum == 1 {
					current = current + (num%limit + 1) + recursive(num%limit)
				} else {
					current = current*lastNum + limit + recursive(num%(limit*lastNum))
				}
				break
			}
		}

		return current
	}

	return recursive(n)
}

// @lc code=end

func Test233() {
	fmt.Println(countDigitOne(88) == 19)
	fmt.Println(countDigitOne(13) == 6)
	fmt.Println(countDigitOne(0) == 0)
}
