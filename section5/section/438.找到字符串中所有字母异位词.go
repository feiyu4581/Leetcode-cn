package section

import "fmt"

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	counter := make(map[byte]int, 26)
	for i := range p {
		counter[p[i]] += 1
	}

	matchCount := 0
	remove := func(char byte) {
		if num, ok := counter[char]; ok {
			counter[char] += 1
			if num == 0 {
				matchCount -= 1
			}
		}
	}

	add := func(char byte) {
		if num, ok := counter[char]; ok {
			counter[char] -= 1
			if num == 1 {
				matchCount += 1
			}
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(p); i++ {
		add(s[i])
	}

	start := 0
	if matchCount == len(counter) {
		res = append(res, start)
	}

	for i := len(p); i < len(s); i++ {
		remove(s[start])
		add(s[i])

		start += 1
		if matchCount == len(counter) {
			res = append(res, start)
		}
	}

	return res
}

// @lc code=end

func Test438() {
	//[0,6]
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

	//[0,1,2]
	fmt.Println(findAnagrams("abab", "ab"))
}
