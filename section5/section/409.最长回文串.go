package section

import "fmt"

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	counter := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counter[s[i]] += 1
	}

	res := 0
	for _, c := range counter {
		res += c
		if c%2 != 0 {
			res -= 1
		}
	}

	if res < len(s) {
		res += 1
	}

	return res
}

// @lc code=end

func Test409() {
	fmt.Println(longestPalindrome("abccccdd") == 7)
	fmt.Println(longestPalindrome("a") == 1)
	fmt.Println(longestPalindrome("aaaaaccc") == 7)
}
