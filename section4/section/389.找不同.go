package section

import "fmt"

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		counter[t[i]-'a'] -= 1
		if counter[t[i]-'a'] < 0 {
			return t[i]
		}
	}

	return 0
}

// @lc code=end

func Test389() {
	fmt.Println(findTheDifference("abcd", "abcde") == 'e')
	fmt.Println(findTheDifference("", "y") == 'y')
}
