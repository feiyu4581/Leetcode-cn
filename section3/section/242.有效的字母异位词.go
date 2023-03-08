package section

import "fmt"

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[byte]int, len(s))
	for index := range s {
		counter[s[index]] += 1
	}

	for index := range t {
		counter[t[index]] -= 1
		if counter[t[index]] < 0 {
			return false
		}
	}

	return true
}

// @lc code=end

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

func Test242() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car") == false)
}
