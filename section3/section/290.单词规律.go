package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	current := 0
	patterCounter := make(map[byte]string, 0)
	wordCounter := make(map[string]byte, 0)
	for current < len(pattern) {
		if _, ok := patterCounter[pattern[current]]; !ok {
			patterCounter[pattern[current]] = words[current]
		}

		if _, ok := wordCounter[words[current]]; !ok {
			wordCounter[words[current]] = pattern[current]
		}

		if patterCounter[pattern[current]] != words[current] || wordCounter[words[current]] != pattern[current] {
			return false
		}

		current += 1
	}

	return true
}

// @lc code=end

func Test290() {
	fmt.Println(wordPattern("abba", "dog dog dog dog") == false)
	fmt.Println(wordPattern("abba", "dog cat cat dog"))
	fmt.Println(wordPattern("abba", "dog cat cat fish") == false)
	fmt.Println(wordPattern("aaaa", "dog cat cat dog") == false)
}
