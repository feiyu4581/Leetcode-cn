package section

import "fmt"

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	counter := make([]int, 26)

	for i := 0; i < len(s); i++ {
		counter[s[i]-'a'] += 1
	}

	for i := 0; i < len(s); i++ {
		if counter[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// @lc code=end

func Test387() {
	fmt.Println(firstUniqChar("leetcode") == 0)
	fmt.Println(firstUniqChar("loveleetcode") == 2)
	fmt.Println(firstUniqChar("aabb") == -1)
}
