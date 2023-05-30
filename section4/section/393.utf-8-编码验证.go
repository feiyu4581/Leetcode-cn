package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */

// @lc code=start
func validUtf8(data []int) bool {
	index := 0

	for index < len(data) {
		if 1<<7&data[index] == 0 {
			index += 1
			continue
		}

		length := 0
		for i := 7; i >= 0; i-- {
			if (1<<i)&data[index] == 0 {
				break
			}
			length += 1
		}

		if length == 1 || length > 4 {
			return false
		}

		for i := 1; i < length; i++ {
			if index+i >= len(data) || data[index+i]&(1<<7+1<<6) != 1<<7 {
				return false
			}
		}

		index += length
	}

	return true
}

// @lc code=end

func Test393() {
	fmt.Println(validUtf8([]int{255}) == false)
	fmt.Println(validUtf8([]int{197, 130, 1}))
	fmt.Println(validUtf8([]int{235, 140, 4}) == false)
	fmt.Println(validUtf8([]int{250, 145, 145, 145, 145}) == false)
}
