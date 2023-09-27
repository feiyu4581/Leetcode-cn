package section

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
var wordGroups = []string{
	"qwertyuiop",
	"asdfghjkl",
	"zxcvbnm",
}

func findWords(words []string) []string {
	wordMaps := make(map[byte]int, 26)
	for i, group := range wordGroups {
		for j := 0; j < len(group); j++ {
			wordMaps[group[j]] = i
		}
	}

	ans := make([]string, 0)
	for wordIndex, word := range words {
		word = strings.ToLower(word)
		num := wordMaps[word[0]]
		isOk := true
		for i := 1; i < len(word); i++ {
			if wordMaps[word[i]] != num {
				isOk = false
				break
			}
		}

		if isOk {
			ans = append(ans, words[wordIndex])
		}
	}

	return ans
}

// @lc code=end

func Test500() {
	// ["Alaska","Dad"]
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))

	// []
	fmt.Println(findWords([]string{"omk"}))

	// ["adsdf","sfd"]
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}
