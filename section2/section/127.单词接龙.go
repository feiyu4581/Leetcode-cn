package section

import "fmt"

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMaps := make(map[string]struct{}, len(wordList))
	for _, word := range wordList {
		wordMaps[word] = struct{}{}
	}

	if _, ok := wordMaps[endWord]; !ok {
		return 0
	}

	currents := []string{beginWord}
	visited := make(map[string]struct{}, len(wordList))
	cache := make(map[string][]string, len(wordList))

	index := 1
	for len(currents) > 0 {
		nextCurrents := make([]string, 0)
		for _, word := range currents {
			if word == endWord {
				return index
			}
		}

		for _, word := range currents {
			for _, nextWord := range getLadderChild(word, wordMaps, cache) {
				if _, ok := visited[nextWord]; ok {
					continue
				}

				nextCurrents = append(nextCurrents, nextWord)
				visited[nextWord] = struct{}{}
			}
		}

		currents = nextCurrents
		index += 1
	}

	return 0
}

func Test127() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}) == 5)
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}) == 0)
}

// @lc code=end
