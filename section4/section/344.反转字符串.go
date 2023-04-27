package section

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString344(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left += 1
		right -= 1
	}
}

// @lc code=end

func Test344() {
	// ['o','l','l','e','h']
	reverseString344([]byte{'h', 'e', 'l', 'l', 'o'})
	// ['h','a','n','n','a','H']
	reverseString344([]byte{'H', 'a', 'n', 'n', 'a', 'h'})
}
