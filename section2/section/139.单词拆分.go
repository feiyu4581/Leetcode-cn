package section

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start

func recursiveWordBreak(s string, index int, maxLength int, wordMap map[string]struct{}, cache map[int]bool) bool {
	if index >= len(s) {
		return true
	}

	if wordOk, ok := cache[index]; ok {
		return wordOk
	}

	for i := 1; i <= maxLength && index+i <= len(s); i++ {
		if _, ok := wordMap[s[index:index+i]]; ok && recursiveWordBreak(s, index+i, maxLength, wordMap, cache) {
			cache[index] = true
			return true
		}
	}

	cache[index] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	if len(wordDict) == 0 {
		return false
	}

	maxLength := math.MinInt
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	return recursiveWordBreak(s, 0, maxLength, wordMap, make(map[int]bool, len(s)))
}

func Test139() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) == false)
}

// @lc code=end
