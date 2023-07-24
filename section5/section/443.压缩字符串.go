package section

import "fmt"

/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

// @lc code=start
func compress(chars []byte) int {
	index := 0
	if len(chars) == 0 {
		return index
	}

	appendChar := func(char byte) {
		chars[index] = char
		index += 1
	}

	appendNum := func(num int) {
		if num == 1 {
			return
		}
		mask := 1
		for mask*10 <= num {
			mask *= 10
		}

		for mask >= 1 {
			digit := num / mask
			appendChar(byte(digit + '0'))
			num %= mask
			mask /= 10
		}
	}

	current := chars[0]
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] == current {
			count += 1
		} else {
			appendChar(current)
			appendNum(count)

			current = chars[i]
			count = 1
		}
	}

	appendChar(current)
	appendNum(count)

	return index
}

// @lc code=end

func Test443() {
	fmt.Println(compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}) == 6)
	fmt.Println(compress([]byte{'a'}) == 1)
	fmt.Println(compress([]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}) == 4)
}
