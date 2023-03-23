package section

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	if len(num) <= 2 {
		return false
	}

	var rangeNumber func(int, int, int) bool
	rangeNumber = func(left, right int, index int) bool {
		sum := left + right
		sumStr := strconv.Itoa(sum)
		if index+len(sumStr) <= len(num) && num[index:index+len(sumStr)] == sumStr {
			// 最终成功标记
			if index+len(sumStr) == len(num) {
				return true
			}

			return rangeNumber(right, sum, index+len(sumStr))
		}

		return false
	}

	mid := len(num) / 2
	for i := 0; i <= mid; i++ {
		for j := i + 1; j <= mid+1; j++ {
			left, _ := strconv.Atoi(num[0 : i+1])
			right, _ := strconv.Atoi(num[i+1 : j+1])
			if rangeNumber(left, right, j+1) {
				return true
			}

			// 数字不能出现以 0 开头的数字，例如 01，如果首字母为0，那么只能以 0 出现一次
			if num[i+1] == '0' {
				break
			}
		}

		// 数字不能出现以 0 开头的数字，例如 01，如果首字母为0，那么只能以 0 出现一次
		if num[0] == '0' {
			break
		}
	}

	return false
}

// @lc code=end

func Test306() {
	fmt.Println(isAdditiveNumber("1023") == false)
	fmt.Println(isAdditiveNumber("0") == false)
	fmt.Println(isAdditiveNumber("111") == false)
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
}
