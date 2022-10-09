package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	res := make([]uint8, 0)
	for columnNumber > 0 {
		num := columnNumber % 26
		if num == 0 {
			columnNumber -= 1
			res = append(res, 'Z')
		} else {
			res = append(res, uint8(columnNumber%26-1)+'A')
		}

		columnNumber = columnNumber / 26
	}

	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left += 1
		right -= 1
	}

	return string(res)
}

func Test168() {
	fmt.Println(convertToTitle(52) == "AZ")
	fmt.Println(convertToTitle(676) == "YZ")
	fmt.Println(convertToTitle(701) == "ZY")
	fmt.Println(convertToTitle(1) == "A")
	fmt.Println(convertToTitle(28) == "AB")
	fmt.Println(convertToTitle(2147483647) == "FXSHRXW")
}

// @lc code=end
