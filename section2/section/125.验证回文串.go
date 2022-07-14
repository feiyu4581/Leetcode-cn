package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	start, end := 0, len(s)-1

	activeChar := func(char uint8) bool {
		if char >= 'a' && char <= 'z' {
			return true
		}

		if char >= '0' && char <= '9' {
			return true
		}

		return false
	}

	s = strings.ToLower(s)
	for start < end {
		left := s[start]
		if !activeChar(left) {
			start++
			continue
		}

		right := s[end]
		if !activeChar(right) {
			end--
			continue
		}

		if left != right {
			return false
		}

		start++
		end--
	}

	return true
}

func Test125() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car") == false)
}

// @lc code=end
