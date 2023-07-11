package section

import "fmt"

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	count := 0

	words := 0
	for i := range s {
		if s[i] == ' ' {
			if words > 0 {
				words = 0
				count += 1
			}
		} else {
			words += 1
		}
	}

	if words > 0 {
		count += 1
	}
	return count
}

// @lc code=end

func Test434() {
	fmt.Println(countSegments("Hello, my name is John") == 5)
}
