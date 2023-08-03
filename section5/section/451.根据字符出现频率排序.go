package section

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start

type CharCount struct {
	Char  byte
	Count int
}

func frequencySort(s string) string {
	countMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
	}

	charCounts := make([]CharCount, 0, len(countMap))
	for k, v := range countMap {
		charCounts = append(charCounts, CharCount{
			Char:  k,
			Count: v,
		})
	}

	sort.Slice(charCounts, func(i, j int) bool {
		return charCounts[i].Count > charCounts[j].Count
	})

	res := make([]byte, 0, len(s))
	for _, charCount := range charCounts {
		for i := 0; i < charCount.Count; i++ {
			res = append(res, charCount.Char)
		}
	}

	return string(res)
}

// @lc code=end

func Test451() {
	// "eert"
	fmt.Println(frequencySort("tree"))
	// "cccaaa"
	fmt.Println(frequencySort("cccaaa"))
	// "bbAa"
	fmt.Println(frequencySort("Aabb"))
}
