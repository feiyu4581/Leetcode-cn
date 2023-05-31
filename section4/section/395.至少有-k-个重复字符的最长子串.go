package section

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a'] += 1
	}

	isMatch := true
	for i := 0; i < len(s); i++ {
		if counter[s[i]-'a'] < k {
			isMatch = false
		}
	}

	if isMatch {
		return len(s)
	}

	start, res := 0, 0
	for start < len(s) {
		if counter[s[start]-'a'] < k {
			start += 1
			continue
		}

		end := start + 1
		for end < len(s) && counter[s[end]-'a'] >= k {
			end += 1
		}

		if end-start > res && end-start >= k {
			if newRes := longestSubstring(s[start:end], k); newRes > res {
				res = newRes
			}
		}

		start = end + 1
	}

	return res

}

// @lc code=end

func Test395() {
	fmt.Println(longestSubstring("aaabb", 3) == 3)
	fmt.Println(longestSubstring("ababbc", 2) == 5)
}
