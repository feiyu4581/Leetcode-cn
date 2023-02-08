package section

import "fmt"

/*
 * @lc app=leetcode.cn id=214 lang=golang
 *
 * [214] 最短回文串
 */

// @lc code=start
func shortestPalindrome(s string) string {

	n := len(s)
	base, mod := 131, 1000000007
	left, right, mul := 0, 0, 1
	best := -1
	for i := 0; i < n; i++ {
		left = (left*base + int(s[i]-'0')) % mod
		right = (right + mul*int(s[i]-'0')) % mod
		if left == right {
			best = i
		}
		mul = mul * base % mod
	}
	add := ""
	if best != n-1 {
		add = s[best+1:]
	}
	b := []byte(add)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b) + s
}

func Test214() {
	fmt.Println(shortestPalindrome("aacecaaa") == "aaacecaaa")
	fmt.Println(shortestPalindrome("abcd") == "dcbabcd")
}

// @lc code=end
