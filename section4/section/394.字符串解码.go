package section

import (
	"bytes"
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start

func decodeRepeatedString(s string, index int) ([]byte, int) {
	var out bytes.Buffer

	for index < len(s) && s[index] != ']' {
		nums := make([]byte, 0)
		for index < len(s) && s[index] >= '0' && s[index] <= '9' {
			nums = append(nums, s[index])
			index += 1
		}

		if index < len(s) && len(nums) > 0 && s[index] == '[' {
			num, _ := strconv.Atoi(string(nums))
			var repeated []byte
			repeated, index = decodeRepeatedString(s, index+1)

			for i := 0; i < num; i++ {
				out.Write(repeated)
			}
		} else {
			out.WriteByte(s[index])
			index += 1
		}
	}

	return out.Bytes(), index + 1
}

func decodeString(s string) string {
	nums, _ := decodeRepeatedString(s, 0)
	return string(nums)
}

// @lc code=end

func Test394() {
	fmt.Println(decodeString("3[a]2[bc]") == "aaabcbc")
	fmt.Println(decodeString("3[a2[c]]") == "accaccacc")
	fmt.Println(decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef")
	fmt.Println(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
}
