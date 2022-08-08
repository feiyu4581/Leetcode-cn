package section

import (
	"fmt"
	"math"
	"strings"
)

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start

func recursiveWordBreak(s string, wordMap map[string]struct{}, index int, maxLength int, cache map[int][][]string) [][]string {
	if index >= len(s) {
		return [][]string{{}}
	}

	if cacheWords, ok := cache[index]; ok {
		return cacheWords
	}

	currents := make([][]string, 0)
	for i := index + 1; i <= index+maxLength && i <= len(s); i++ {
		if _, ok := wordMap[s[index:i]]; ok {
			for _, subWords := range recursiveWordBreak(s, wordMap, i, maxLength, cache) {
				words := make([]string, 0, len(subWords)+1)
				words = append(words, s[index:i])
				words = append(words, subWords...)
				currents = append(currents, words)
			}
		}
	}

	cache[index] = currents
	return currents
}

func wordBreak(s string, wordDict []string) []string {
	if s == "" {
		return nil
	}

	if len(wordDict) == 0 {
		return nil
	}

	maxLength := math.MinInt
	wordMap := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}

	words := make([]string, 0)
	for _, wordList := range recursiveWordBreak(s, wordMap, 0, maxLength, make(map[int][][]string, 0)) {
		words = append(words, strings.Join(wordList, " "))
	}

	return words
}

func Test140() {
	// ["cats and dog","cat sand dog"]
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	// ["pine apple pen apple","pineapple pen apple","pine applepen apple"]
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	// []
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

// @lc code=end
