package section

import "fmt"

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
// 元音字母：'a'、'e'、'i'、'o'、'u'

var vowelMaps = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		for left < right {
			if _, ok := vowelMaps[bytes[left]]; ok {
				break
			}
			left += 1
		}

		for left < right {
			if _, ok := vowelMaps[bytes[right]]; ok {
				break
			}
			right -= 1
		}

		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left += 1
			right -= 1
		}
	}

	return string(bytes)
}

// @lc code=end

func Test345() {
	fmt.Println(reverseVowels("hello") == "holle")
	fmt.Println(reverseVowels("leetcode") == "leotcede")
}
