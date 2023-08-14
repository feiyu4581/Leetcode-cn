package section

import "fmt"

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 */

// @lc code=start
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}

	checkSubString := func(end int) bool {
		start := 0
		for i := end; i < len(s); i++ {
			if start == end {
				start = 0
			}

			if s[i] != s[start] {
				return false
			}
			start++
		}

		return start == end
	}

	for i := 1; i <= len(s)/2; i++ {
		if s[i] == s[0] && checkSubString(i) {
			return true
		}
	}

	return false
}

// @lc code=end

func Test459() {
	fmt.Println(repeatedSubstringPattern("aabaaba") == false)
	fmt.Println(repeatedSubstringPattern("abab") == true)
	fmt.Println(repeatedSubstringPattern("aba") == false)
	fmt.Println(repeatedSubstringPattern("abcabcabcabc") == true)
	fmt.Println(repeatedSubstringPattern("ababab") == true)
	fmt.Println(repeatedSubstringPattern("abababab") == true)
}
